## Overview

Ever play a game where you choose a page in Wikipedia and see how many
hops it takes to get to a different page in Wikipedia? For example, how
many hops does it take to get from Cheese to Antartica. This program 
takes a start and end subject and determines how many hops it takes in 
Wikipedia to get there. 

Using a CLI the user enters a starting page/subject and an end goal for
the program. They are returned back with the path it took to get there. 

## CLI Guide

This program uses the Cobra Cli Library offering for Golang: https://github.com/spf13/cobra

To get the path printed out from start -> finish including the total number of hops 
the CLI example is as follows: 
```
wikiConnections --start Ballet --goal Tennis
Tennis
Gymnastics
Ballet
Number of hops is:  3
```
There are two search alogirthms used in this program: Depth First Search (DFS) and
Breadth First Search (BFS). The program will use BFS as the default search alogrithm. 
If a user would like to use DFS they will have to add the following flag to their command: 
`wikiConnections --start Ballet --goal Tennis --algorithm DFS`

The user can also dictate how many pages we will allow the algorithm to visit before 
giving up on finding the path between start -> goal. The default value of maximum pages
visited is 20 however this can also be changed via the CLI with the following `-p --pageLimit`
flag: `wikiConnections --start Ballet --goal Tennis --pageLimit 15`