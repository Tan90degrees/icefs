/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-04-20 16:45:13
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-21 02:48:22
 * @FilePath: /icefs/src/lowlevel/client/include/icefsThriftClientPool.hpp
 * @Description:
 */
#ifndef ICEFS_THRIFT_CLIENT_POOL_HPP
#define ICEFS_THRIFT_CLIENT_POOL_HPP

#include <mutex>
#include <queue>

#include "icefsClient.hpp"

#define MAX_THRIFT_CLIENT_NUM 100

struct icefsThriftConn {
  std::shared_ptr<apache::thrift::transport::TTransport> thriftSocket;
  std::shared_ptr<apache::thrift::transport::TTransport> thriftTransport;
  std::shared_ptr<apache::thrift::protocol::TProtocol> thriftProtocol;
  icefsthrift::IcefsThriftClient *thriftClient;
};

class icefsThriftConnPool {
 private:
  std::mutex poolLock;
  std::queue<icefsThriftConn *> clientPool;
  uint64_t maxClient = MAX_THRIFT_CLIENT_NUM;
  IcefsClientConfig clientConfig;

 public:
  icefsThriftConnPool(const IcefsClientConfig *config) {
    this->clientConfig = *config;
  }
  ~icefsThriftConnPool() {
    poolLock.lock();
    icefsThriftConn *conn = nullptr;
    while (!clientPool.empty()) {
      conn = clientPool.front();
      clientPool.pop();
      conn->thriftTransport->close();
      delete conn;
      conn = nullptr;
    }
  }

  icefsThriftConn *GetIcefsThriftConn(void) {
    icefsThriftConn *conn;

    poolLock.lock();
    if (!clientPool.empty()) {
      conn = clientPool.front();
      clientPool.pop();
      poolLock.unlock();
    } else {
      poolLock.unlock();
      conn = new icefsThriftConn();
      conn->thriftSocket = std::make_shared<apache::thrift::transport::TSocket>(
          clientConfig.serverAddress, clientConfig.port);
      conn->thriftTransport =
          std::make_shared<apache::thrift::transport::TBufferedTransport>(
              conn->thriftSocket);
      conn->thriftProtocol =
          std::make_shared<apache::thrift::protocol::TBinaryProtocol>(
              conn->thriftTransport);
      conn->thriftClient =
          new icefsthrift::IcefsThriftClient(conn->thriftProtocol);
      conn->thriftTransport->open();
    }
    return conn;
  }

  void PutIcefsThriftConn(icefsThriftConn *conn) {
    poolLock.lock();
    if (clientPool.size() < maxClient) {
      clientPool.push(conn);
      poolLock.unlock();
    } else {
      poolLock.unlock();
      conn->thriftTransport->close();
      delete conn;
    }
    return;
  }
};

#endif