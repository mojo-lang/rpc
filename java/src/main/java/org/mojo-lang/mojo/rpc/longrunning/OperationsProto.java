// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/rpc/longrunning/operations.proto

package org.mojo-lang.mojo.rpc.longrunning;

public final class OperationsProto {
  private OperationsProto() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_mojo_rpc_longrunning_ListOperationsRequest_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_mojo_rpc_longrunning_ListOperationsRequest_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_mojo_rpc_longrunning_ListOperationsResponse_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_mojo_rpc_longrunning_ListOperationsResponse_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_mojo_rpc_longrunning_GetOperationRequest_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_mojo_rpc_longrunning_GetOperationRequest_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_mojo_rpc_longrunning_DeleteOperationRequest_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_mojo_rpc_longrunning_DeleteOperationRequest_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_mojo_rpc_longrunning_CancelOperationRequest_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_mojo_rpc_longrunning_CancelOperationRequest_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_mojo_rpc_longrunning_WaitOperationRequest_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_mojo_rpc_longrunning_WaitOperationRequest_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n%mojo/rpc/longrunning/operations.proto\022" +
      "\024mojo.rpc.longrunning\032\024mojo/core/null.pr" +
      "oto\032\024mojo/core/time.proto\032$mojo/rpc/long" +
      "running/operation.proto\"o\n\025ListOperation" +
      "sRequest\022\016\n\006parent\030\001 \001(\t\022\016\n\006filter\030\004 \001(\t" +
      "\022\022\n\tpage_size\030\320\017 \001(\005\022\023\n\npage_token\030\321\017 \001(" +
      "\t\022\r\n\004skip\030\322\017 \001(\005\"}\n\026ListOperationsRespon" +
      "se\0223\n\noperations\030\001 \003(\0132\037.mojo.rpc.longru" +
      "nning.Operation\022\024\n\013total_count\030\320\017 \001(\005\022\030\n" +
      "\017next_page_token\030\321\017 \001(\t\"#\n\023GetOperationR" +
      "equest\022\014\n\004name\030\001 \001(\t\"&\n\026DeleteOperationR" +
      "equest\022\014\n\004name\030\001 \001(\t\"&\n\026CancelOperationR" +
      "equest\022\014\n\004name\030\001 \001(\t\"J\n\024WaitOperationReq" +
      "uest\022\014\n\004name\030\001 \001(\t\022$\n\007timeout\030\002 \001(\0132\023.mo" +
      "jo.core.Duration2\334\003\n\nOperations\022l\n\017list_" +
      "operations\022+.mojo.rpc.longrunning.ListOp" +
      "erationsRequest\032,.mojo.rpc.longrunning.L" +
      "istOperationsResponse\022[\n\rget_operation\022)" +
      ".mojo.rpc.longrunning.GetOperationReques" +
      "t\032\037.mojo.rpc.longrunning.Operation\022Q\n\020de" +
      "lete_operation\022,.mojo.rpc.longrunning.De" +
      "leteOperationRequest\032\017.mojo.core.Null\022Q\n" +
      "\020cancel_operation\022,.mojo.rpc.longrunning" +
      ".CancelOperationRequest\032\017.mojo.core.Null" +
      "\022]\n\016wait_operation\022*.mojo.rpc.longrunnin" +
      "g.WaitOperationRequest\032\037.mojo.rpc.longru" +
      "nning.OperationBy\n\"org.mojo-lang.mojo.rp" +
      "c.longrunningB\017OperationsProtoP\001Z@github" +
      ".com/mojo-lang/rpc/go/pkg/mojo/rpc/longr" +
      "unning;longrunningb\006proto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
          org.mojolang.mojo.core.NullProto.getDescriptor(),
          org.mojolang.mojo.core.TimeProto.getDescriptor(),
          org.mojo-lang.mojo.rpc.longrunning.OperationProto.getDescriptor(),
        });
    internal_static_mojo_rpc_longrunning_ListOperationsRequest_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_mojo_rpc_longrunning_ListOperationsRequest_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_mojo_rpc_longrunning_ListOperationsRequest_descriptor,
        new java.lang.String[] { "Parent", "Filter", "PageSize", "PageToken", "Skip", });
    internal_static_mojo_rpc_longrunning_ListOperationsResponse_descriptor =
      getDescriptor().getMessageTypes().get(1);
    internal_static_mojo_rpc_longrunning_ListOperationsResponse_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_mojo_rpc_longrunning_ListOperationsResponse_descriptor,
        new java.lang.String[] { "Operations", "TotalCount", "NextPageToken", });
    internal_static_mojo_rpc_longrunning_GetOperationRequest_descriptor =
      getDescriptor().getMessageTypes().get(2);
    internal_static_mojo_rpc_longrunning_GetOperationRequest_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_mojo_rpc_longrunning_GetOperationRequest_descriptor,
        new java.lang.String[] { "Name", });
    internal_static_mojo_rpc_longrunning_DeleteOperationRequest_descriptor =
      getDescriptor().getMessageTypes().get(3);
    internal_static_mojo_rpc_longrunning_DeleteOperationRequest_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_mojo_rpc_longrunning_DeleteOperationRequest_descriptor,
        new java.lang.String[] { "Name", });
    internal_static_mojo_rpc_longrunning_CancelOperationRequest_descriptor =
      getDescriptor().getMessageTypes().get(4);
    internal_static_mojo_rpc_longrunning_CancelOperationRequest_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_mojo_rpc_longrunning_CancelOperationRequest_descriptor,
        new java.lang.String[] { "Name", });
    internal_static_mojo_rpc_longrunning_WaitOperationRequest_descriptor =
      getDescriptor().getMessageTypes().get(5);
    internal_static_mojo_rpc_longrunning_WaitOperationRequest_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_mojo_rpc_longrunning_WaitOperationRequest_descriptor,
        new java.lang.String[] { "Name", "Timeout", });
    org.mojolang.mojo.core.NullProto.getDescriptor();
    org.mojolang.mojo.core.TimeProto.getDescriptor();
    org.mojo-lang.mojo.rpc.longrunning.OperationProto.getDescriptor();
  }

  // @@protoc_insertion_point(outer_class_scope)
}