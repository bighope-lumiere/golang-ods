# Open Data Structures with Golang

## これはなに

Go 言語に入門したので、練習がてら [Open Data Structures](http://opendatastructures.org/) のデータ構造を Go 言語で実装してみようと思います。

## 実装したいデータ構造

まえがきで特に重要だと書いてあるデータ構造は最低限実装したいな...

- 第2章: ArrayStack, ArrayQueue, ArrayDeque
- 第3章: SLList, DLList
- 第5章: ChainedHashTable
- 第6章: BinaryTree, BinarySearchTree
- 第9章: RedBlackTree
- 第10章: BinaryHeap
- 第11章: MergeSort, QuickSort
- 第12章: 幅優先探索, 深さ優先探索

## List インターフェースとは

List インターフェースは、値の列 `x[0], ..., x[n-1]` と、その列に対する以下の操作からなる

1. `size()`: リストの長さ n を返す
2. `get(i)`: `x[i]` を返す
3. `set(i, x)`: `x[i]` を `x` にする
4. `add(i, x)`: `x` を `i` 番目として追加し、 `x[i], ..., x[n-1]` を後ろにずらす
5. `remove(i)`: `x[i]` を削除し、 `x[i+1], ..., x[n-1]` を前にずらす

## USet インターフェースとは

USet (Unordered Set) インターフェースは集合のようなインターフェースであり、互いに異なる n 個の要素が含まれる。要素の並び順は決まっていない。

USet インターフェースに実行できる操作は以下

1. `size()`: 集合の要素数 `n` を返す
2. `add(x)`: 要素 `x` が集合に入っていなければ集合に追加する
3. `remove(x)`: 集合から `x` を削除する
4. `find(x)`: 集合に `x` が入っていればそれを見つける

## SSet インターフェースとは

SSet (Sorted Set) インターフェースは順序付けされた要素の集まりを表現する（全順序集合）。基本的に操作は USet と同じで、 `find(x)` だけ異なる。

- `find(x)`: 順序付けされた集合から `x` の位置を特定する -> つまり y>=x なる最小の要素 y を見つける
