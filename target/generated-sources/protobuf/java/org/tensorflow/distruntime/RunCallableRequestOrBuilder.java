// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: tensorflow/core/protobuf/master.proto

package org.tensorflow.distruntime;

public interface RunCallableRequestOrBuilder extends
    // @@protoc_insertion_point(interface_extends:tensorflow.RunCallableRequest)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <pre>
   * REQUIRED: session_handle must be returned by a CreateSession call
   * to the same master service.
   * </pre>
   *
   * <code>string session_handle = 1;</code>
   */
  java.lang.String getSessionHandle();
  /**
   * <pre>
   * REQUIRED: session_handle must be returned by a CreateSession call
   * to the same master service.
   * </pre>
   *
   * <code>string session_handle = 1;</code>
   */
  com.google.protobuf.ByteString
      getSessionHandleBytes();

  /**
   * <pre>
   * REQUIRED: handle must be returned by a MakeCallable call to the same
   * master service.
   * </pre>
   *
   * <code>int64 handle = 2;</code>
   */
  long getHandle();

  /**
   * <pre>
   * Values of the tensors passed as arguments to the callable, in the order
   * defined in the CallableOptions.feed field passed to MakeCallable.
   * </pre>
   *
   * <code>repeated .tensorflow.TensorProto feed = 3;</code>
   */
  java.util.List<org.tensorflow.framework.TensorProto> 
      getFeedList();
  /**
   * <pre>
   * Values of the tensors passed as arguments to the callable, in the order
   * defined in the CallableOptions.feed field passed to MakeCallable.
   * </pre>
   *
   * <code>repeated .tensorflow.TensorProto feed = 3;</code>
   */
  org.tensorflow.framework.TensorProto getFeed(int index);
  /**
   * <pre>
   * Values of the tensors passed as arguments to the callable, in the order
   * defined in the CallableOptions.feed field passed to MakeCallable.
   * </pre>
   *
   * <code>repeated .tensorflow.TensorProto feed = 3;</code>
   */
  int getFeedCount();
  /**
   * <pre>
   * Values of the tensors passed as arguments to the callable, in the order
   * defined in the CallableOptions.feed field passed to MakeCallable.
   * </pre>
   *
   * <code>repeated .tensorflow.TensorProto feed = 3;</code>
   */
  java.util.List<? extends org.tensorflow.framework.TensorProtoOrBuilder> 
      getFeedOrBuilderList();
  /**
   * <pre>
   * Values of the tensors passed as arguments to the callable, in the order
   * defined in the CallableOptions.feed field passed to MakeCallable.
   * </pre>
   *
   * <code>repeated .tensorflow.TensorProto feed = 3;</code>
   */
  org.tensorflow.framework.TensorProtoOrBuilder getFeedOrBuilder(
      int index);
}
