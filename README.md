Compile command: ```go build -o {executable_name} *.go``` or to run immediately ```go run main.go expr-queue.go expression.go```

A command line mental math game. You will be shown a math expression and must evaluate it.

Every 3 seconds, a new expression will be added into a queue, and once that queue exceeds the size of 12, it will be game over.

To win, you can answer fast to empty the queue or keep the queue size less than 12 for 60 seconds.

There are 3 levels to the game (1 - easy, 2 - medium, 3 - hard). The higher the level, the greater the operands.

Sample video: (sorry for the slow mental math skills in advance :sweat_smile:)

https://user-images.githubusercontent.com/73256760/177717531-65657700-accb-43ed-9075-175d844b310b.mov


The [C++ implementation](https://github.com/nealarch01/MentalMathCLIGame) of this project
