package com.example.allinone2.selfrpc;

import com.example.allinone2.selfrpc.annotify.LzyService;

import java.io.IOException;
import java.lang.reflect.InvocationTargetException;
import java.net.ServerSocket;

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


@LzyService
class TestServiceImpl implements TestService {
    @Override
    public String foo(String arg1, String arg2) {
        return arg1 + " and " + arg2;
    }
}
