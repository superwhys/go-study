#!/usr/bin/env python
# -*- coding: UTF-8 -*-
#
# @File    ：client.py
# @IDE     ：PyCharm 
# @Author  ：SuperYong
# @Date    ：2021/10/11 18:30 
# @Summary : this is the summary
from loguru import logger
from grpc import insecure_channel
from pb import pb_pb2_grpc, pb_pb2


def run():
    # 注意(gRPC Python Team): .close()方法在channel上是可用的。
    # 并且应该在with语句不符合代码需求的情况下使用。
    with insecure_channel('localhost:8972') as channel:
        stub = pb_pb2_grpc.GreeterStub(channel)
        response = stub.SayHello(pb_pb2.HelloRequest(name='q1mi'))
    logger.info("Greeter client received: {}!".format(response.message))


if __name__ == '__main__':
    run()
