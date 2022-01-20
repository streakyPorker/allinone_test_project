package com.example.allinone2.reflect;

import java.lang.reflect.InvocationHandler;
import java.lang.reflect.Method;
import java.lang.reflect.Proxy;
import java.util.Arrays;

interface TestInterface {
    /**
     * asd
     *
     * @param str asd
     * @return asd
     */
    String foo(String str);
}

/**
 * @author lzy
 */
public class ReflectTest {
    public static void main(String[] args) {
        TestInterface testInterface = generateImpl(TestInterface.class);
        System.out.println(testInterface.foo("asd"));
    }

    static void wrap(Object... obj){

    }

    static <T> T generateImpl(Class<T> intf) {
        System.out.println("info is " + intf.getTypeName());

        Method[] methods = intf.getMethods();
        InvocationHandler handler = new InvocationHandler() {
            @Override
            public Object invoke(Object proxy, Method method, Object[] args) throws Throwable {
                System.out.println(method);
                System.out.println("method name:"+method.getName()+", args:"+ Arrays.toString(args)+", return type:"+method.getReturnType());
                if(method.getReturnType().equals(String.class)){
                    return "string val";
                }
                return method.getReturnType().newInstance();
            }
        };
        T hello = (T) Proxy.newProxyInstance(
            intf.getClassLoader(),
            new Class[] { intf },
            handler);
        return  hello;
    }
}
