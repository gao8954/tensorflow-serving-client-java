// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: tensorflow/core/framework/device_attributes.proto

package org.tensorflow.framework;

public interface LocalLinksOrBuilder extends
    // @@protoc_insertion_point(interface_extends:tensorflow.LocalLinks)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <code>repeated .tensorflow.InterconnectLink link = 1;</code>
   */
  java.util.List<org.tensorflow.framework.InterconnectLink> 
      getLinkList();
  /**
   * <code>repeated .tensorflow.InterconnectLink link = 1;</code>
   */
  org.tensorflow.framework.InterconnectLink getLink(int index);
  /**
   * <code>repeated .tensorflow.InterconnectLink link = 1;</code>
   */
  int getLinkCount();
  /**
   * <code>repeated .tensorflow.InterconnectLink link = 1;</code>
   */
  java.util.List<? extends org.tensorflow.framework.InterconnectLinkOrBuilder> 
      getLinkOrBuilderList();
  /**
   * <code>repeated .tensorflow.InterconnectLink link = 1;</code>
   */
  org.tensorflow.framework.InterconnectLinkOrBuilder getLinkOrBuilder(
      int index);
}
