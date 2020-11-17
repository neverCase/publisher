package runner

import (
	"fmt"
	"github.com/nevercase/publisher/pkg/interfaces"
	"github.com/nevercase/publisher/pkg/types"
	"k8s.io/klog"
)

const (
	StepOperatorWasNotExisted = "err: the specific interfaces.StepOperator step-name:%s was not existed"
)

type Runner struct {
	Name          string                    `json:"name" protobuf:"bytes,1,opt,name=name"`
	Hostname      string                    `json:"hostname" protobuf:"bytes,2,opt,name=hostname"`
	Namespace     types.Namespace           `json:"namespace" protobuf:"bytes,3,opt,name=namespace"`
	GroupName     types.GroupName           `json:"groupName" protobuf:"bytes,4,opt,name=groupName"`
	StepOperators []interfaces.StepOperator `json:"stepOperators" protobuf:"bytes,5,opt,name=stepOperators"`
	// StreamOutput was a chan<- string which was used to transfer exec outputs by the stream.
	StreamOutput chan<- string `json:"streamOutput"`
}

func (r *Runner) Register() (res types.RunnerInfo, err error) {
	steps := make([]types.Step, 0)
	for _, v := range r.StepOperators {
		steps = append(steps, *v.Step())
	}
	res = types.RunnerInfo{
		Name:       r.Name,
		Hostname:   r.Hostname,
		Namespace:  r.Namespace,
		GroupName:  r.GroupName,
		RunnerType: types.RunnerTypeServer,
		Steps:      steps,
	}
	return res, nil
}

func (r *Runner) Run(s *types.Step) (err error) {
	exist := false
	for _, v := range r.StepOperators {
		if v.Step().Name == s.Name {
			v.Update(s)
			exist = true
			v.Prepare()
			res, err := v.Run(r.StreamOutput)
			if err != nil {
				klog.V(2).Info(err)
				// todo report Run error
				return err
			}
			// todo update Step information use DeepCopy
			_ = res
		}
	}
	if !exist {
		// todo report Run started failed
		return fmt.Errorf(StepOperatorWasNotExisted, s.Name)
	}
	return nil
}

func (r *Runner) Update(s *types.Step) (err error) {
	exist := false
	for _, v := range r.StepOperators {
		if v.Step().Name == s.Name {
			exist = true
			v.Step().Envs = s.Envs
		}
	}
	if !exist {
		return fmt.Errorf(StepOperatorWasNotExisted, s.Name)
	}
	return nil
}

func (r *Runner) Step(s *types.Step) (*types.Step, error) {
	for _, v := range r.StepOperators {
		if v.Step().Name == s.Name {
			return v.Step(), nil
		}
	}
	return nil, fmt.Errorf(StepOperatorWasNotExisted, s.Name)
}
