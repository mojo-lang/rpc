// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/rpc/longrunning/operations.proto

package org.mojo-lang.mojo.rpc.longrunning;

public interface WaitOperationRequestOrBuilder extends
    // @@protoc_insertion_point(interface_extends:mojo.rpc.longrunning.WaitOperationRequest)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <code>string name = 1;</code>
   * @return The name.
   */
  java.lang.String getName();
  /**
   * <code>string name = 1;</code>
   * @return The bytes for name.
   */
  com.google.protobuf.ByteString
      getNameBytes();

  /**
   * <code>.mojo.core.Duration timeout = 2;</code>
   * @return Whether the timeout field is set.
   */
  boolean hasTimeout();
  /**
   * <code>.mojo.core.Duration timeout = 2;</code>
   * @return The timeout.
   */
  org.mojolang.mojo.core.Duration getTimeout();
  /**
   * <code>.mojo.core.Duration timeout = 2;</code>
   */
  org.mojolang.mojo.core.DurationOrBuilder getTimeoutOrBuilder();
}
