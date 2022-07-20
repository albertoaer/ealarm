package core

type ReentrantCommand interface {
	//Launch the reiterative command, must pass only one value to the channel,
	//true -> continue execution, false -> end alarm iteration
	Launch(next chan bool)
}

type ReentrantCommandBuilder func() ReentrantCommand

type multiCmd []ReentrantCommand

func (m multiCmd) Launch(next chan bool) {
	for i, cmd := range m {
		cmd.Launch(next)
		if i < len(m)-1 {
			v := <-next
			if !v {
				next <- v
				break
			}
		}
	}
}

func ManyCommands(cmds ...ReentrantCommand) ReentrantCommand {
	return multiCmd(cmds)
}

func ManyCommandsBuilder(builders ...ReentrantCommandBuilder) ReentrantCommandBuilder {
	return func() ReentrantCommand {
		cmds := make([]ReentrantCommand, len(builders))
		for i := range builders {
			cmds[i] = builders[i]()
		}
		return ManyCommands(cmds...)
	}
}
