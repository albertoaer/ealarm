package core

type ReentrantCommand interface {
	//Launch the reiterative command, must pass only one value to the channel,
	//true -> continue execution, false -> end alarm iteration
	Launch(next chan bool)
}

type multiCmd []ReentrantCommand

func (m multiCmd) Launch(next chan bool) {
	for _, cmd := range m {
		cmd.Launch(next)
		v := <-next
		next <- v
		if !v {
			break
		}
	}
}

func ManyCommands(cmds ...ReentrantCommand) ReentrantCommand {
	return multiCmd(cmds)
}
