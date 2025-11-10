### Go语言实现的 IK 控制
<img src="fk_test.webp" alt="fk_test"></img><br>
正向动力学：根据骨骼状态计算最终位置<br>
<img src="ik_test.webp" alt="ik_test"/></img><br>
反向动力学：根据最终位置计算骨骼状态<br>
<img src="ik.webp" alt="ik"/></img><br>
大概原理
### 视频演示
https://github.com/user-attachments/assets/2a125004-2d6e-4202-9a12-a3e3cbf0431b
红色线垂直于父骨骼，蓝色箭头表示骨骼，绿色扇形标识骨骼允许旋转的范围，视频中的骨骼始终在跟随光标