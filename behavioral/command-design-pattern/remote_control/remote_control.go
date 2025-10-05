package remotecontrol

import "github.com/roydevashish/golang-design-patterns/command/command"

type RemoteControl struct {
	cmd command.Command
}

func NewRemoteControl(cmd command.Command) *RemoteControl {
	return &RemoteControl{
		cmd: cmd,
	}
}

func (r *RemoteControl) SetCommand(cmd command.Command) {
	r.cmd = cmd
}

func (r *RemoteControl) PressButton() {
	r.cmd.Execute()
}
