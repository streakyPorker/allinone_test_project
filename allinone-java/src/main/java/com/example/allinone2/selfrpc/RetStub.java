package com.example.allinone2.selfrpc;

import java.io.Serializable;

import lombok.Data;

@Data
public class RetStub extends Stub implements Serializable {
    Object[] args;
    Class[] argTypes;
    Class<?> retType;
    Object retVal;
 }
