<!--
 * @Author: 27
 * @LastEditors: 27
 * @Date: 2022-04-26 14:27:51
 * @LastEditTime: 2022-04-26 14:34:49
 * @FilePath: /graphs-Rishabh-Mishra/Explain_doc.md
 * @description: type some description
-->
# 解释说明

## reference
作者自己写的解释算法的文章
> [Dijkstra Algorithm in Golang](https://medium.com/@rishabhmishra131/golang-dijkstra-algorithm-7bf2722ba0c8

## Algorithm
- Create a Distance Map Storing Distance from source Node to Each Nodes, Initialize all the valuse to Infinity
- Create a Vertex Map to mark vertexes as visited and non visited
- For source Node mark distance as zero (weights should be non negative)
- Enqueue the source Node in min Priority queue , Priority Queue is based on distance
- Dequeue the Node from queue and visit all unvisited neighbours , calculate tentative distances through current node, if tentative distance is less than current assign distance , update the distance and assign the smaller one, enqueue that smallest distance vertex in priority queue
- To Get shortest Path in step 5 store previous node name
- Repeat this untill priority queue is empty
