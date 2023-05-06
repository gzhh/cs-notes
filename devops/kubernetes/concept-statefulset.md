## StatefulSet

StatefulSet 与 ReplicaSet 或 ReplicationController 的对比

- StatefulSet 在 ReplicaSet 和 ReplicationController 的功能上做了一些升级
- StatefulSet 保证了 pod 在重新调度后保留它们的标志和状态，让你更方便的扩容、缩容
