package com.example.allinone2.selfrpc;

import java.io.Closeable;
import java.io.IOException;
import java.io.ObjectInputStream;
import java.io.ObjectOutputStream;
import java.lang.reflect.InvocationTargetException;
import java.lang.reflect.Method;
import java.net.ServerSocket;
import java.net.Socket;
import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;
import java.util.concurrent.Executors;
import java.util.concurrent.ThreadPoolExecutor;

/**
 * @author lzy
 */

public class P2pRpcServiceCenter implements Closeable {

    private final ThreadPoolExecutor tpe = (ThreadPoolExecutor)Executors.newFixedThreadPool(10);

    private final ServerSocket ss;
    private final Map<String, Object> serviceMap = new ConcurrentHashMap<>();

    public P2pRpcServiceCenter(ServerSocket ss) {
        this.ss = ss;
    }

    public void serve() {
        while (true) {
            try {
                Socket socket = ss.accept();
                tpe.execute(generateServeTask(socket));
            } catch (IOException e) {
                e.printStackTrace();
            }
        }
    }

    public <T> boolean registerService(Class<T> clazz, T impl) {
        return serviceMap.putIfAbsent(clazz.getName(), impl) == null;
    }

    @Override
    public void close() {

    }

    private Runnable generateServeTask(Socket socket) {
        return () -> {
            System.out.println("get a req:"+socket);
            ObjectInputStream ois = null;
            ObjectOutputStream oos = null;
            try {
                ois = new ObjectInputStream(socket.getInputStream());
                CallStub callStub = (CallStub)ois.readObject();
                if (serviceMap.containsKey(callStub.getServiceName())) {
                    oos = new ObjectOutputStream(socket.getOutputStream());
                    Object impl = serviceMap.get(callStub.getServiceName());
                    Method method = impl.getClass().getMethod(callStub.getMethodName(), callStub.argTypes);
                    RetStub retStub = new RetStub();
                    retStub.methodName = callStub.methodName;
                    retStub.retVal = method.invoke(impl, callStub.args);
                    retStub.retType = method.getReturnType();
                    retStub.args = callStub.getArgs();
                    retStub.argTypes = callStub.getArgTypes();
                    oos.writeObject(retStub);
                    ois.close();
                    oos.close();
                }
            } catch (IOException | ClassNotFoundException | NoSuchMethodException | InvocationTargetException | IllegalAccessException e) {
                e.printStackTrace();
            }
        };
    }
}
