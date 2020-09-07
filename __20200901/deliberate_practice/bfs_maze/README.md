刻意练习，超容易理解的 Golang BFS 走迷宫

## 介绍

编码思维训练，不训练就秀逗了🐶🐶。

这会是一个系列。

## The Maze（迷宫）

下面是一个 `6` 行 `5` 列的二维阵列，用来表示迷宫。

* `0` 代表可以通行
* `1` 表示不能通行

我|是|迷|宫|宫
-|-|-|-|-
0|1|0|0|0
0|0|0|1|0
0|1|0|1|0
1|1|1|0|0
0|1|0|0|1
0|1|0|0|0

我们是使用一个二维 `slice` 来表示它。

```go
maze := [][]int {
  {0, 1, 0, 0, 0},
  {0, 0, 0, 1, 0},
  {0, 1, 0, 1, 0},
  {1, 1, 1, 0, 0},
  {0, 1, 0, 0, 1},
  {0, 1, 0, 0, 0},
}
```

既然是走迷宫，必定有`入口`、`出口`、和可以探索出口的`方向`。

这里设定：
* `入口` 是二维数组坐标 `(0, 0)`
* `出口` 是二维数组坐标 `(5, 4)`
* `方向` 是 `上`、`下`、`左`、`右`

### 如何从入口走到出口呢？

想象自己被扔在一个未知的点，我们允许在这一点的上、下、左、右`四`个方向去探索。

通过`1`步，能走到的格子是`4`个。

```
  1
1 0 1
  1
```

通过`2`步，能走到的格子是`8`个。

```
    2
  2 1 2
2 1 0 1 2
  2 1 2
    2
```

通过`3`步，能走到的格子是`12`个。

```
      3
    3 2 3
  3 2 1 2 3
3 2 1 0 1 2 3
  3 2 1 2 3
    3 2 3
      3
```

通过`n`步，能走到的格子是`n*4(方向)`个。

当前这一步，或者说当前这一点的状态就有如下：
* 已经发现但还没有探索（排队）
* 已经探索
* 发都还没有发现

`1`(4个方向) 探索完了才轮到 `2`，是一个很自然的排队过程。

一层一层的往外递进，确保每到一个点都是一个最短的路径到的这个点。

**应用：**

注意：**写程序的时候上、左、下、右。逆时针 90°，90°的转。可能会带来一些好处。**

1. 空投到了一个点`(0, 0)`，0 步被走到。其它的各自是未知的，迷宫长啥样也是未知的。
  * 状态：已经发现但还没有探索（`排队`->也就是首先将 `(0, 0)`）
2. 开始探索 (0, 0)--> 上，左，右，下
  * 上，左 -> 出界❌
  * 下(1,0) -> 可以，我们标一个 `1`，表示可以一步走到这个点
    * 放入队列-->`已经发现但还没有探索（排队）`
  * 右 -> 墙❌
3. 开始探索 (1, 0)--> 上，左，右，下
  * 上 -> 已探索❌
  * 左 -> 出界❌
  * 下 -> 可以，我们标一个 `2`，表示可以两步走到这个点
    * 放入队列(2, 0)-->`已经发现但还没有探索（排队）`
  * 右 -> 可以，我们标一个 `2`，表示可以两步走到这个点
    * 放入队列(1, 1)-->`已经发现但还没有探索（排队）`
4. 开始探索 (2, 0)
  ......发现并入队
5. 开始探索 (1, 1)
  ......发现并入队
6. 开始探索 (x, y)......

**结束条件：**
* 已经走到终点
* 队列为空

**基本编码流程：**

1. 先把迷宫读进来
2. 走迷宫（walk）
  * `start` ---> `end`
  * 从某一点`如：(0, 0)`走到某一点`如：(5, 4)`
3. 空降到 `start`，我们要维护另外一个 `2` 维的 `slice(steps)`
  * 里面每一格代表从 `start` 走了多少步才到这一格
  * `steps`
    * 这个很重要，最后的路径就是用这个建立出来的
4. 一个格子完成探索的需要做两件事儿
  * 格子放入起点到达它的步数
  * 将发现的格子放入队列

### 开始编码

```go
package main

import "fmt"

type point struct {
	x int
	y int
}

// 下一个点
func (p point) add(d point) point {
	return point{
		p.x + d.x,
		p.y + d.y,
	}
}

func (p point) at(grid [][]int) (int, bool) {
	// 上下越界
	if p.x < 0 || p.x >= len(grid) {
		return 0, false
	}
	// 左右越界
	if p.y < 0 || p.y >= len(grid[p.x]) {
		return 0, false
	}
	// 返回 grid[p.x][p.y] 的好处 --> 撞墙，已经探索等等统统放外层判断
	return grid[p.x][p.y], true
}

func walk(maze [][]int, start point, end point) [][]int {
	// 维护一个与 maze 相同的 Steps Slice
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[0]))
	}
	// 定义一个队列，并将起点入队
	// start: 已经发现但还没有探索（排队）
	Q := []point{start}

	// 定义 上，左，下，右 四个方向
	dirs := []point{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

	// 经典写法，队列为空，走迷宫结束
	for len(Q) > 0 {
		cur := Q[0] // 要探索的点
		Q = Q[1:]   // 出队

		// 四个方向探索
		for _, d := range dirs {
			// 下一点
			next := cur.add(d)
			// val：用来判断在 maze 是否是墙(1)
			// ok: 用来判断在 maze 是否越界
			val, ok := next.at(maze)
			// 有墙不能探索，越界不能探索
			if val == 1 || !ok {
				continue
			}
			// val：用来判断点在 steps 是否是值，有值表明这个位置在 `maze` 中已经探索过了
			// ok: 用来判断在 steps 是否越界
			val, ok = next.at(steps)
			if val != 0 || !ok {
				continue
			}
			// 下一探索点不能是起点
			if next == start {
				continue
			}
			curSteps, _ := cur.at(steps)
			// 格子放入起点到达它的步数
			steps[next.x][next.y] = curSteps + 1
			// 将发现的格子放入队列--> 已经发现但还没有探索（排队）
			Q = append(Q, next)
		}
	}
	return steps
}

func main() {
	// 迷宫
	maze := [][]int{
		{0, 1, 0, 0, 0},
		{0, 0, 0, 1, 0},
		{0, 1, 0, 1, 0},
		{1, 1, 1, 0, 0},
		{0, 1, 0, 0, 1},
		{0, 1, 0, 0, 0},
	}
	// 入口
	start := point{0, 0}
	// 出口
	end := point{len(maze) - 1, len(maze[0]) - 1}

	steps := walk(maze, start, end)
	for _, row := range steps {
		for _, val := range row {
			// 3位对齐，打印结果
			fmt.Printf("%3d", val)
		}
		fmt.Println()
	}
}

```

输|出|结|果|果
-|-|-|-|-
0|0|`4`|`5`|`6`
`1`|`2`|`3`|0|`7`
2|0|4|0|`8`
0|0|0|`10`|`9`
0|0|12|`11`|0
0|0|13|`12`|`13`