package com.example.selfrpc;

import java.io.IOException;
import java.io.ObjectInputStream;
import java.io.ObjectOutputStream;
import java.lang.reflect.InvocationTargetException;
import java.lang.reflect.Method;
import java.net.ServerSocket;
import java.net.Socket;

/**
 * @author lzy
 */
public class SelfProvider {
    public static void main(String[] args) throws IOException, ClassNotFoundException, NoSuchMethodException, InvocationTargetException, IllegalAccessException {
        ServerSocket ss = new ServerSocket(9999);
        P2pRpcServiceCenter p2pRpcServiceCenter = new P2pRpcServiceCenter(ss);
        p2pRpcServiceCenter.registerService(TestService.class, new TestServiceImpl());
        p2pRpcServiceCenter.serve();

    }
}

class TestServiceImpl implements TestService {
    @Override
    public String foo(String arg1, String arg2) {
        return arg1 + " and " + arg2;
    }
}
