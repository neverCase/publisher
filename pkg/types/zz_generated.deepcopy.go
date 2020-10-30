// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by deepcopy-gen. DO NOT EDIT.

package types

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Group) DeepCopyInto(out *Group) {
	*out = *in
	if in.Tasks != nil {
		in, out := &in.Tasks, &out.Tasks
		*out = make([]Task, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Runners != nil {
		in, out := &in.Runners, &out.Runners
		*out = make([]RunnerInfo, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Group.
func (in *Group) DeepCopy() *Group {
	if in == nil {
		return nil
	}
	out := new(Group)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ListGroupNameRequest) DeepCopyInto(out *ListGroupNameRequest) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ListGroupNameRequest.
func (in *ListGroupNameRequest) DeepCopy() *ListGroupNameRequest {
	if in == nil {
		return nil
	}
	out := new(ListGroupNameRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ListGroupNameResponse) DeepCopyInto(out *ListGroupNameResponse) {
	*out = *in
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ListGroupNameResponse.
func (in *ListGroupNameResponse) DeepCopy() *ListGroupNameResponse {
	if in == nil {
		return nil
	}
	out := new(ListGroupNameResponse)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ListNamespaceRequest) DeepCopyInto(out *ListNamespaceRequest) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ListNamespaceRequest.
func (in *ListNamespaceRequest) DeepCopy() *ListNamespaceRequest {
	if in == nil {
		return nil
	}
	out := new(ListNamespaceRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ListNamespaceResponse) DeepCopyInto(out *ListNamespaceResponse) {
	*out = *in
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ListNamespaceResponse.
func (in *ListNamespaceResponse) DeepCopy() *ListNamespaceResponse {
	if in == nil {
		return nil
	}
	out := new(ListNamespaceResponse)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ListTaskRequest) DeepCopyInto(out *ListTaskRequest) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ListTaskRequest.
func (in *ListTaskRequest) DeepCopy() *ListTaskRequest {
	if in == nil {
		return nil
	}
	out := new(ListTaskRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ListTaskResponse) DeepCopyInto(out *ListTaskResponse) {
	*out = *in
	if in.Tasks != nil {
		in, out := &in.Tasks, &out.Tasks
		*out = make([]Task, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ListTaskResponse.
func (in *ListTaskResponse) DeepCopy() *ListTaskResponse {
	if in == nil {
		return nil
	}
	out := new(ListTaskResponse)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogStream) DeepCopyInto(out *LogStream) {
	*out = *in
	if in.Output != nil {
		in, out := &in.Output, &out.Output
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogStream.
func (in *LogStream) DeepCopy() *LogStream {
	if in == nil {
		return nil
	}
	out := new(LogStream)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegisterRunnerRequest) DeepCopyInto(out *RegisterRunnerRequest) {
	*out = *in
	in.RunnerInfo.DeepCopyInto(&out.RunnerInfo)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegisterRunnerRequest.
func (in *RegisterRunnerRequest) DeepCopy() *RegisterRunnerRequest {
	if in == nil {
		return nil
	}
	out := new(RegisterRunnerRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegisterRunnerResponse) DeepCopyInto(out *RegisterRunnerResponse) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegisterRunnerResponse.
func (in *RegisterRunnerResponse) DeepCopy() *RegisterRunnerResponse {
	if in == nil {
		return nil
	}
	out := new(RegisterRunnerResponse)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Request) DeepCopyInto(out *Request) {
	*out = *in
	out.Type = in.Type
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Request.
func (in *Request) DeepCopy() *Request {
	if in == nil {
		return nil
	}
	out := new(Request)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Response) DeepCopyInto(out *Response) {
	*out = *in
	out.Type = in.Type
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Response.
func (in *Response) DeepCopy() *Response {
	if in == nil {
		return nil
	}
	out := new(Response)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Result) DeepCopyInto(out *Result) {
	*out = *in
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Result.
func (in *Result) DeepCopy() *Result {
	if in == nil {
		return nil
	}
	out := new(Result)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RunnerInfo) DeepCopyInto(out *RunnerInfo) {
	*out = *in
	if in.Steps != nil {
		in, out := &in.Steps, &out.Steps
		*out = make([]Step, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RunnerInfo.
func (in *RunnerInfo) DeepCopy() *RunnerInfo {
	if in == nil {
		return nil
	}
	out := new(RunnerInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Step) DeepCopyInto(out *Step) {
	*out = *in
	if in.Envs != nil {
		in, out := &in.Envs, &out.Envs
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Output != nil {
		in, out := &in.Output, &out.Output
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.UploadFiles != nil {
		in, out := &in.UploadFiles, &out.UploadFiles
		*out = make([]UploadFile, len(*in))
		copy(*out, *in)
	}
	if in.WriteFiles != nil {
		in, out := &in.WriteFiles, &out.WriteFiles
		*out = make([]WriteFile, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Messages != nil {
		in, out := &in.Messages, &out.Messages
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Step.
func (in *Step) DeepCopy() *Step {
	if in == nil {
		return nil
	}
	out := new(Step)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Task) DeepCopyInto(out *Task) {
	*out = *in
	if in.Runners != nil {
		in, out := &in.Runners, &out.Runners
		*out = make(map[string]RunnerInfo, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Task.
func (in *Task) DeepCopy() *Task {
	if in == nil {
		return nil
	}
	out := new(Task)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Type) DeepCopyInto(out *Type) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Type.
func (in *Type) DeepCopy() *Type {
	if in == nil {
		return nil
	}
	out := new(Type)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UploadFile) DeepCopyInto(out *UploadFile) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UploadFile.
func (in *UploadFile) DeepCopy() *UploadFile {
	if in == nil {
		return nil
	}
	out := new(UploadFile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WriteFile) DeepCopyInto(out *WriteFile) {
	*out = *in
	if in.Content != nil {
		in, out := &in.Content, &out.Content
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WriteFile.
func (in *WriteFile) DeepCopy() *WriteFile {
	if in == nil {
		return nil
	}
	out := new(WriteFile)
	in.DeepCopyInto(out)
	return out
}
