// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: tensorflow/core/protobuf/saver.proto

package org.tensorflow.util;

public interface SaverDefOrBuilder extends
    // @@protoc_insertion_point(interface_extends:tensorflow.SaverDef)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <pre>
   * The name of the tensor in which to specify the filename when saving or
   * restoring a model checkpoint.
   * </pre>
   *
   * <code>string filename_tensor_name = 1;</code>
   */
  java.lang.String getFilenameTensorName();
  /**
   * <pre>
   * The name of the tensor in which to specify the filename when saving or
   * restoring a model checkpoint.
   * </pre>
   *
   * <code>string filename_tensor_name = 1;</code>
   */
  com.google.protobuf.ByteString
      getFilenameTensorNameBytes();

  /**
   * <pre>
   * The operation to run when saving a model checkpoint.
   * </pre>
   *
   * <code>string save_tensor_name = 2;</code>
   */
  java.lang.String getSaveTensorName();
  /**
   * <pre>
   * The operation to run when saving a model checkpoint.
   * </pre>
   *
   * <code>string save_tensor_name = 2;</code>
   */
  com.google.protobuf.ByteString
      getSaveTensorNameBytes();

  /**
   * <pre>
   * The operation to run when restoring a model checkpoint.
   * </pre>
   *
   * <code>string restore_op_name = 3;</code>
   */
  java.lang.String getRestoreOpName();
  /**
   * <pre>
   * The operation to run when restoring a model checkpoint.
   * </pre>
   *
   * <code>string restore_op_name = 3;</code>
   */
  com.google.protobuf.ByteString
      getRestoreOpNameBytes();

  /**
   * <pre>
   * Maximum number of checkpoints to keep.  If 0, no checkpoints are deleted.
   * </pre>
   *
   * <code>int32 max_to_keep = 4;</code>
   */
  int getMaxToKeep();

  /**
   * <pre>
   * Shard the save files, one per device that has Variable nodes.
   * </pre>
   *
   * <code>bool sharded = 5;</code>
   */
  boolean getSharded();

  /**
   * <pre>
   * How often to keep an additional checkpoint. If not specified, only the last
   * "max_to_keep" checkpoints are kept; if specified, in addition to keeping
   * the last "max_to_keep" checkpoints, an additional checkpoint will be kept
   * for every n hours of training.
   * </pre>
   *
   * <code>float keep_checkpoint_every_n_hours = 6;</code>
   */
  float getKeepCheckpointEveryNHours();

  /**
   * <code>.tensorflow.SaverDef.CheckpointFormatVersion version = 7;</code>
   */
  int getVersionValue();
  /**
   * <code>.tensorflow.SaverDef.CheckpointFormatVersion version = 7;</code>
   */
  org.tensorflow.util.SaverDef.CheckpointFormatVersion getVersion();
}
