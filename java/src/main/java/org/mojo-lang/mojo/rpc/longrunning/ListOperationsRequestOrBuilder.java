// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/rpc/longrunning/operations.proto

package org.mojo-lang.mojo.rpc.longrunning;

public interface ListOperationsRequestOrBuilder extends
    // @@protoc_insertion_point(interface_extends:mojo.rpc.longrunning.ListOperationsRequest)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <code>string parent = 1;</code>
   * @return The parent.
   */
  java.lang.String getParent();
  /**
   * <code>string parent = 1;</code>
   * @return The bytes for parent.
   */
  com.google.protobuf.ByteString
      getParentBytes();

  /**
   * <code>string filter = 4;</code>
   * @return The filter.
   */
  java.lang.String getFilter();
  /**
   * <code>string filter = 4;</code>
   * @return The bytes for filter.
   */
  com.google.protobuf.ByteString
      getFilterBytes();

  /**
   * <code>int32 page_size = 2000;</code>
   * @return The pageSize.
   */
  int getPageSize();

  /**
   * <code>string page_token = 2001;</code>
   * @return The pageToken.
   */
  java.lang.String getPageToken();
  /**
   * <code>string page_token = 2001;</code>
   * @return The bytes for pageToken.
   */
  com.google.protobuf.ByteString
      getPageTokenBytes();

  /**
   * <code>int32 skip = 2002;</code>
   * @return The skip.
   */
  int getSkip();
}
