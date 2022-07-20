# Ealarm

Command line tool to run cyclic commands, including a default alarm ui

# Features

- Alarm configuration
  - duration
  - display message
  - track to play
  - number of times to play, infinite for n < 0
  - cyclic action to run
  - compound actions of other actions, syntax: "Action1&Action2"
- Presets file
  - actions: User defined command line actions
  - profiles: Collection of predefined Ealarm configurations

# Why?

Looking for an option in the default desktop alarm to loop it a custom amount of time I did not found it, so I decided to create my own alarm