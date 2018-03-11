## 六个图

| 0  | 1 | 2 | 3 |
| -------- | -----: | :----: | ----: | 
| ui-原型 | use case -用例 | robustness-diagram 分析图 | 序列图 | 
|        | 领域图 | 修整过的 领域图 | 类图 | 



## noun-verb-noun 模式

约束： 除了名词不能直接跟名词交互  其他都可以

- 边界类 n  
- 控制类 v  
- 实体类 n

序列图天生就是完全的 n-v-n 模式
v 表对象间的消息流

> Remember that controllers are only occasionally real control objects; they are more
typically logical software functions.

> Keep in mind that both boundary objects and entity objects are nouns, and that controllers
are verbs (i.e., an action performed on an object). As such, it makes sense that the
controllers (the actions) will become methods on the boundary and entity classes.

> Remember That a Robustness Diagram Is an “Object Picture” of a Use Case
A robustness diagram is an “object picture” of a use case, whose purpose is to force refinement
of both use case text and the object model. Robustness diagrams tie use cases to objects (and
to the GUI).