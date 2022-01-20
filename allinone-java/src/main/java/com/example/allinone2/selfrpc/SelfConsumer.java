package com.example.allinone2.selfrpc;

import java.io.IOException;
import java.io.ObjectInputStream;
import java.io.ObjectOutputStream;
import java.lang.reflect.InvocationHandler;
import java.lang.reflect.Method;
import java.lang.reflect.Proxy;
import java.net.InetAddress;
import java.net.Socket;

public class SelfConsumer {

    public static void main(String[] args) throws IOException {
        TestService testService = getRpcService();
        System.out.println(testService.foo("arg1", "arg2"));
    }

    public static <T> T getRpcService(Class<T> service) throws IOException {
        Socket socket = new Socket(InetAddress.getLocalHost(), 9999);
        CallStub callStub = new CallStub();
        return null;
    }

    public static TestService getRpcService() throws IOException {
        Socket socket = new Socket(InetAddress.getLocalHost(), 9999);
        CallStub callStub = new CallStub();
        callStub.setServiceName(TestService.class.getName());

        InvocationHandler handler = new InvocationHandler() {
            @Override
            public Object invoke(Object proxy, Method method, Object[] args) throws Throwable {
                ObjectOutputStream oos = new ObjectOutputStream(socket.getOutputStream());
                callStub.methodName = method.getName();
                callStub.argTypes = method.getParameterTypes();
                callStub.args = args;
                callStub.retType = method.getReturnType();
                oos.writeObject(callStub);
                ObjectInputStream ois = new ObjectInputStream(socket.getInputStream());
                RetStub retStub = (RetStub) ois.readObject();
                for (int i = 0; i < args.length; i++) {
                    args[i] = retStub.args[i];
                }
                return retStub.retVal;
            }
        };
        return (TestService) Proxy.newProxyInstance(
                TestService.class.getClassLoader(),
                new Class[]{TestService.class},
                handler);
    }


}
