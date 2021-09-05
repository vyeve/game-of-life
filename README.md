# Game of Life

## Start

Implementation the game of life.

To start run next command:

```shell
go run main.go
```
## Patterns

There are some implemented patterns:

|  #|  pattern | 
| ------------ | ------------ |
|   1| [glider ](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life#/media/File:Game_of_life_animated_glider.gif "glider ") |
|  2 | [pulsar](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life#/media/File:Game_of_life_pulsar.gif "pulsar")  |
|  3 |[Middle- weight spaceship (MWSS)](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life#/media/File:Animated_Mwss.gif "Middle- weight spaceship (MWSS)")   |
|   4|[Gosper glider gun](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life#/media/File:Game_of_life_glider_gun.svg "Gosper glider gun")   |
|  5 |  [Penta-decathlon](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life#/media/File:I-Column.gif "Penta-decathlon") |
| 6  |  Random generated pattern (40x20) |

To run selected pattern, please run command:

```shell
# run pulsar pattern
go run -p 2
# run penta-decathlon pattern
go run -p 5
```
