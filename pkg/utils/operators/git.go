package operators

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/Shanghai-Lunara/publisher/pkg/interfaces"
	"github.com/Shanghai-Lunara/publisher/pkg/types"
	"k8s.io/klog/v2"
)

func NewGit(gitDir string, branchName string) interfaces.StepOperator {
	envs := make(map[string]string, 0)
	envs[types.PublisherProjectDir] = gitDir
	envs[types.PublisherGitBranch] = branchName
	return &git{
		output: make(chan<- string, 4096),
		step: &types.Step{
			Id:             0,
			Name:           "Git-Operator",
			Phase:          types.StepPending,
			Policy:         types.StepPolicyAuto,
			Available:      types.StepAvailableEnable,
			Envs:           envs,
			Output:         make([]string, 0),
			SharingData:    make(map[string]string, 0),
			SharingSetting: false,
		},
	}
}

type git struct {
	output chan<- string
	step   *types.Step
}

func (g *git) Step() *types.Step {
	return g.step
}

func (g *git) Update(s *types.Step) {
	g.step = s.DeepCopy()
}

func (g *git) Prepare() {

}

func (g *git) Run(output chan<- string) (res []string, err error) {
	g.output = output
	g.step.Phase = types.StepRunning
	var out []byte
	if out, err = g.cd(); err != nil {
		klog.V(2).Info(err)
		g.step.Phase = types.StepFailed
		return res, err
	}
	if out, err = g.revert(); err != nil {
		klog.V(2).Info(err)
		g.step.Phase = types.StepFailed
		return res, err
	}
	if out, err = g.checkout(); err != nil {
		klog.V(2).Info(err)
		g.step.Phase = types.StepFailed
		return res, err
	}
	if out, err = g.pull(); err != nil {
		klog.V(2).Info(err)
		g.step.Phase = types.StepFailed
		return res, err
	}
	res = append(res, string(out))
	g.step.Phase = types.StepSucceeded
	return res, nil
}

func (g *git) cd() (res []byte, err error) {
	commands := fmt.Sprintf("cd %s", g.step.Envs[types.PublisherProjectDir])
	return DefaultExec(commands)
}

func (g *git) branch() (res string, err error) {
	commands := fmt.Sprintf("cd %s && git branch -a | grep '*'", g.step.Envs[types.PublisherProjectDir])
	t, err := DefaultExec(commands)
	if err != nil {
		klog.V(2).Info(err)
		return res, err
	}
	activeMatched, err := regexp.Match(`\*`, t)
	if err != nil {
		return res, fmt.Errorf("git regexp.Match active target-name:%s err:%v\n", g.step.Envs[types.PublisherGitBranch], err)
	}
	if activeMatched == false {
		return res, fmt.Errorf("git regexp.Match active target-name:%s failed", g.step.Envs[types.PublisherGitBranch])
	}
	var name string
	name = strings.Replace(string(t), " ", "", -1)
	name = strings.Replace(name, "*", "", -1)
	// such as `* test\n`
	name = strings.Replace(name, "\n", "", -1)
	return name, nil
}

func (g *git) fetchAll() (res []byte, err error) {
	commands := fmt.Sprintf("cd %s && git fetch --all && git fetch -p", g.step.Envs[types.PublisherProjectDir])
	return ExecWithStreamOutput(commands, g.output)
}

func (g *git) revert() (res []byte, err error) {
	commands := fmt.Sprintf("cd %s && git add --all && git checkout -f && git reset --hard", g.step.Envs[types.PublisherProjectDir])
	return ExecWithStreamOutput(commands, g.output)
}

func (g *git) checkout() (res []byte, err error) {
	commands := fmt.Sprintf("cd %s && git checkout -B %s --track remotes/origin/%s",
		g.step.Envs[types.PublisherProjectDir], g.step.Envs[types.PublisherGitBranch], g.step.Envs[types.PublisherGitBranch])
	klog.Info("git checkout commands:", commands)
	return ExecWithStreamOutput(commands, g.output)
}

func (g *git) pull() (res []byte, err error) {
	commands := fmt.Sprintf("cd %s && git pull", g.step.Envs[types.PublisherProjectDir])
	return ExecWithStreamOutput(commands, g.output)
}

func (g *git) push() (res []byte, err error) {
	commands := fmt.Sprintf("cd %s && git push", g.step.Envs[types.PublisherProjectDir])
	return ExecWithStreamOutput(commands, g.output)
}
