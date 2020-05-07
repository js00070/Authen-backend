# 技术介绍

## Swarm分布式存储原理

Swarm是一种点对点数据共享网络所组成的分布式文件存储系统，其中文件通过其内容的哈希来寻址。与Bittorrent类似，可以同时从多个节点获取数据，只要单个节点承载分发数据，它就可以随处被访问。这种方法可以在不必依靠托管任何类型服务器的情况下分发数据，数据可访问性与位置无关。Swarm通过以太坊来激励网络中的其他节点自己复制和存储数据，能保证数据的高可用，即使原节点宕机也可以从其他备份了此文件的节点来获取到文件内容。

Swarm中不存在删除/移除功能。数据一旦上传，就被永久多副本存储在了节点网络中，文件内容无法被修改。

Swarm定义了三个关键概念：

- **Chunk**：大小有限 (最大4K) 的数据块，Swarm中存储和检索的基本单位。网络层只识别chunk，没有文件概念。

- **Reference**：文件的唯一标识符，允许客户端检索和访问内容。对于未加密内容，文件reference是数据的加密哈希，并作为其内容地址。该哈希长度为32字节，序列化为64个十六进制字节。如果是加密文件，则包含两个等长的部分：前32个字节是内容地址，后32个字节是解密密钥（共64字节），序列化为128个十六进制字节。

- **Manifest**：描述文件集合的数据结构。Manifest指定路径和相应的内容哈希，以允许基于URL的内容检索。Manifest也可以映射到文件系统目录树 (directory tree)，该目录树允上传和下载目录。最后，manifest也可以被视作是索引，因此它可以用于实现简单的键值存储或数据库索引。

![img](https://gblobscdn.gitbook.com/assets%2F-LpXMJ2UtB5ZJjI4I8qn%2F-M2XFVd1cEKBE97hEPBg%2F-M2XOynbhJ224VknVLkW%2Fdapp-page.svg?alt=media&token=bd6c9309-21dd-49cf-94fd-8da12e0cbb6b)

上图展示了一个利用Swarm来获取网页资源文件的流程：

1. 客户端发送root-hash到Swarm节点
2. Swarm节点从节点网络中收集相应的manifest并返回
3. 客户端把manifest作为索引，从Swarm网络中获取到所有的资源文件
4. 得到了所有资源文件，渲染页面

## 利用Swarm实现电子证据未篡改证明

### 流程

1. 用户使用app客户端，进行录音采集或是照片、图片拍摄，产生的电子证据文件实时进行加密存储在手机本地，客户无法使用常规手段对其进行非破坏性的修改（无法导出后使用PhotoShop或是录音、视频剪辑工具进行篡改）

2. app在文件采集完成后会把此文件的元信息（包括创建时间、文件内容哈希值等）立刻发送给服务器。服务器接收到元信息后存储在本地数据库，同时会将元信息同步至本地Swarm节点中，由本地Swarm节点完成文件元信息在公网的多副本同步
3. 用户为了防止电子证据所在的设备被物理损毁，可以将文件本体也上传至服务器，服务器会验证文件完整性和正确性，确保文件未被篡改，并将文件保留在服务器中作为备份
4. 当用户涉及到司法事务，需要证明电子证据文件未被篡改时，只需要从Swarm网络中下载到文件元信息，对照自己本地的文件或是服务器备份的文件，对文件内容哈希值进行比对，只要哈希值一致，就可以确认电子证据未被篡改，具有法律效益