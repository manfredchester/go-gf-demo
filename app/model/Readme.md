# 数据模型

数据模型负责维护所有的数据操作，特别是对数据库的访问控制。数据模型往往是被service调用，不推荐通过控制器直接访问数据模型。

数据模型的代码位于/app/model。