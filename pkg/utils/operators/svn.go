package operators

import (
	"fmt"
	"github.com/nevercase/publisher/pkg/interfaces"
	"github.com/nevercase/publisher/pkg/types"
	"k8s.io/klog/v2"
)

func NewSvn(host string, port int, username, password, remoteDir, workDir string) interfaces.StepOperator {
	envs := make(map[string]string, 0)
	envs[types.PublisherSvnHost] = host
	envs[types.PublisherSvnPort] = fmt.Sprintf("%d", port)
	envs[types.PublisherSvnUsername] = username
	envs[types.PublisherSvnPassword] = password
	envs[types.PublisherSvnRemoteDir] = remoteDir
	envs[types.PublisherSvnWorkDir] = workDir
	envs[types.PublisherSvnCommitMessage] = "svn commit"
	envs[types.PublisherSvnCommand] = SvnCommandWaiting
	return &svn{
		step: &types.Step{
			Id:       0,
			Name:     "SVN-Operator",
			Phase:    types.StepPending,
			Policy:   types.StepPolicyAuto,
			Envs:     envs,
			Messages: make([]string, 0),
			Output:   make([]string, 0),
		},
	}
}

type svn struct {
	output chan<- string
	step   *types.Step
}

func (s *svn) Step() *types.Step {
	return s.step
}

func (s *svn) Update(step *types.Step) {
	s.step = step.DeepCopy()
}

func (s *svn) Prepare() {
	s.step.Messages = make([]string, 0)
}

const (
	SvnCommandWaiting    = "pulling and waiting"
	SvnCommandCommitting = "adding and committing"
	SvnCommandCD         = "cd"
	SvnCommandRevertAll  = "revert all"
	SvnCommandRemoveAll  = "remove all"
	SvnCommandCheckout   = "checkout"
	SvnCommandAddAll     = "add all"
	SvnCommandCommit     = "commit"
)

func (s *svn) AppendMessage(action string) {
	s.step.Messages = append(s.step.Messages, types.StepMessage(s.step.Name, action))
}

func (s *svn) Run(output chan<- string) (res []string, err error) {
	s.output = output
	s.step.Phase = types.StepRunning
	var out []byte
	switch s.step.Envs[types.PublisherSvnCommand] {
	case SvnCommandWaiting:
		s.AppendMessage(SvnCommandWaiting)
		s.AppendMessage(SvnCommandCD)
		if out, err = s.cd(); err != nil {
			klog.V(2).Info(err)
			s.step.Phase = types.StepFailed
			return res, err
		}
		res = append(res, string(out))
		s.AppendMessage(SvnCommandRevertAll)
		if out, err = s.revertAll(); err != nil {
			klog.V(2).Info(err)
			s.step.Phase = types.StepFailed
			return res, err
		}
		res = append(res, string(out))
		s.AppendMessage(SvnCommandRemoveAll)
		if out, err = s.removeAll(); err != nil {
			klog.V(2).Info(err)
			s.step.Phase = types.StepFailed
			return res, err
		}
		res = append(res, string(out))
		s.AppendMessage(SvnCommandCheckout)
		if out, err = s.checkout(); err != nil {
			klog.V(2).Info(err)
			s.step.Phase = types.StepFailed
			return res, err
		}
		res = append(res, string(out))
	case SvnCommandCommitting:
		s.AppendMessage(SvnCommandCommitting)
		s.AppendMessage(SvnCommandAddAll)
		if out, err = s.addAll(); err != nil {
			klog.V(2).Info(err)
			s.step.Phase = types.StepFailed
			return res, err
		}
		res = append(res, string(out))
		s.AppendMessage(SvnCommandCommit)
		if out, err = s.commit(); err != nil {
			klog.V(2).Info(err)
			s.step.Phase = types.StepFailed
			return res, err
		}
		res = append(res, string(out))
	}
	s.step.Phase = types.StepSucceeded
	return res, nil
}

func (s *svn) cd() (res []byte, err error) {
	commands := fmt.Sprintf("cd %s", s.step.Envs[types.PublisherSvnWorkDir])
	return DefaultExec(commands)
}

const svnUrl = "svn://%s@%s:%s/%s"

func (s *svn) checkout() (res []byte, err error) {
	commands := fmt.Sprintf("cd %s && svn --username %s --password %s checkout %s",
		s.step.Envs[types.PublisherSvnWorkDir],
		s.step.Envs[types.PublisherSvnUsername],
		s.step.Envs[types.PublisherSvnPassword],
		fmt.Sprintf(svnUrl,
			s.step.Envs[types.PublisherSvnUsername],
			s.step.Envs[types.PublisherSvnHost],
			s.step.Envs[types.PublisherSvnPort],
			s.step.Envs[types.PublisherSvnRemoteDir]),
	)
	return ExecWithStreamOutput(commands, s.output)
}

func (s *svn) addAll() (res []byte, err error) {
	commands := fmt.Sprintf("cd %s && svn --username %s --password %s status | grep ? | awk '{print $2}' | xargs svn --username %s --password %s add",
		fmt.Sprintf("%s/%s", s.step.Envs[types.PublisherSvnWorkDir], s.step.Envs[types.PublisherSvnRemoteDir]),
		s.step.Envs[types.PublisherSvnUsername],
		s.step.Envs[types.PublisherSvnPassword],
		s.step.Envs[types.PublisherSvnUsername],
		s.step.Envs[types.PublisherSvnPassword],
	)
	return ExecWithStreamOutput(commands, s.output)
}

func (s *svn) revertAll() (res []byte, err error) {
	commands := fmt.Sprintf("cd %s && svn --username %s --password %s status | awk '{print $2}' | xargs svn --username %s --password %s revert --depth infinity",
		fmt.Sprintf("%s/%s", s.step.Envs[types.PublisherSvnWorkDir], s.step.Envs[types.PublisherSvnRemoteDir]),
		s.step.Envs[types.PublisherSvnUsername],
		s.step.Envs[types.PublisherSvnPassword],
		s.step.Envs[types.PublisherSvnUsername],
		s.step.Envs[types.PublisherSvnPassword],
	)
	return ExecWithStreamOutput(commands, s.output)
}

func (s *svn) removeAll() (res []byte, err error) {
	commands := fmt.Sprintf("cd %s && svn --username %s --password %s status | grep ? | awk '{print $2}' | xargs rm -rf",
		fmt.Sprintf("%s/%s", s.step.Envs[types.PublisherSvnWorkDir], s.step.Envs[types.PublisherSvnRemoteDir]),
		s.step.Envs[types.PublisherSvnUsername],
		s.step.Envs[types.PublisherSvnPassword],
	)
	return ExecWithStreamOutput(commands, s.output)
}

func (s *svn) commit() (res []byte, err error) {
	commands := fmt.Sprintf("cd %s && svn --username %s --password %s commit --message \"${%s} - committed by ${%s}@publisher\"",
		fmt.Sprintf("%s/%s", s.step.Envs[types.PublisherSvnWorkDir], s.step.Envs[types.PublisherSvnRemoteDir]),
		s.step.Envs[types.PublisherSvnUsername],
		s.step.Envs[types.PublisherSvnPassword],
		s.step.Envs[types.PublisherSvnCommitMessage],
		s.step.Envs[types.PublisherSvnUsername],
	)
	return ExecWithStreamOutput(commands, s.output)
}
