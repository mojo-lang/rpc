// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/rpc/longrunning/operation_service.proto

package org.mojo-lang.mojo.rpc.longrunning;

/**
 * Protobuf type {@code mojo.rpc.longrunning.Operations}
 */
public final class Operations extends
    com.google.protobuf.GeneratedMessageV3 implements
    // @@protoc_insertion_point(message_implements:mojo.rpc.longrunning.Operations)
    OperationsOrBuilder {
private static final long serialVersionUID = 0L;
  // Use Operations.newBuilder() to construct.
  private Operations(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
    super(builder);
  }
  private Operations() {
    vals_ = java.util.Collections.emptyList();
  }

  @java.lang.Override
  @SuppressWarnings({"unused"})
  protected java.lang.Object newInstance(
      UnusedPrivateParameter unused) {
    return new Operations();
  }

  @java.lang.Override
  public final com.google.protobuf.UnknownFieldSet
  getUnknownFields() {
    return this.unknownFields;
  }
  private Operations(
      com.google.protobuf.CodedInputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    this();
    if (extensionRegistry == null) {
      throw new java.lang.NullPointerException();
    }
    int mutable_bitField0_ = 0;
    com.google.protobuf.UnknownFieldSet.Builder unknownFields =
        com.google.protobuf.UnknownFieldSet.newBuilder();
    try {
      boolean done = false;
      while (!done) {
        int tag = input.readTag();
        switch (tag) {
          case 0:
            done = true;
            break;
          case 10: {
            if (!((mutable_bitField0_ & 0x00000001) != 0)) {
              vals_ = new java.util.ArrayList<org.mojo-lang.mojo.rpc.longrunning.Operation>();
              mutable_bitField0_ |= 0x00000001;
            }
            vals_.add(
                input.readMessage(org.mojo-lang.mojo.rpc.longrunning.Operation.parser(), extensionRegistry));
            break;
          }
          default: {
            if (!parseUnknownField(
                input, unknownFields, extensionRegistry, tag)) {
              done = true;
            }
            break;
          }
        }
      }
    } catch (com.google.protobuf.InvalidProtocolBufferException e) {
      throw e.setUnfinishedMessage(this);
    } catch (java.io.IOException e) {
      throw new com.google.protobuf.InvalidProtocolBufferException(
          e).setUnfinishedMessage(this);
    } finally {
      if (((mutable_bitField0_ & 0x00000001) != 0)) {
        vals_ = java.util.Collections.unmodifiableList(vals_);
      }
      this.unknownFields = unknownFields.build();
      makeExtensionsImmutable();
    }
  }
  public static final com.google.protobuf.Descriptors.Descriptor
      getDescriptor() {
    return org.mojo-lang.mojo.rpc.longrunning.OperationServiceProto.internal_static_mojo_rpc_longrunning_Operations_descriptor;
  }

  @java.lang.Override
  protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internalGetFieldAccessorTable() {
    return org.mojo-lang.mojo.rpc.longrunning.OperationServiceProto.internal_static_mojo_rpc_longrunning_Operations_fieldAccessorTable
        .ensureFieldAccessorsInitialized(
            org.mojo-lang.mojo.rpc.longrunning.Operations.class, org.mojo-lang.mojo.rpc.longrunning.Operations.Builder.class);
  }

  public static final int VALS_FIELD_NUMBER = 1;
  private java.util.List<org.mojo-lang.mojo.rpc.longrunning.Operation> vals_;
  /**
   * <code>repeated .mojo.rpc.longrunning.Operation vals = 1;</code>
   */
  @java.lang.Override
  public java.util.List<org.mojo-lang.mojo.rpc.longrunning.Operation> getValsList() {
    return vals_;
  }
  /**
   * <code>repeated .mojo.rpc.longrunning.Operation vals = 1;</code>
   */
  @java.lang.Override
  public java.util.List<? extends org.mojo-lang.mojo.rpc.longrunning.OperationOrBuilder> 
      getValsOrBuilderList() {
    return vals_;
  }
  /**
   * <code>repeated .mojo.rpc.longrunning.Operation vals = 1;</code>
   */
  @java.lang.Override
  public int getValsCount() {
    return vals_.size();
  }
  /**
   * <code>repeated .mojo.rpc.longrunning.Operation vals = 1;</code>
   */
  @java.lang.Override
  public org.mojo-lang.mojo.rpc.longrunning.Operation getVals(int index) {
    return vals_.get(index);
  }
  /**
   * <code>repeated .mojo.rpc.longrunning.Operation vals = 1;</code>
   */
  @java.lang.Override
  public org.mojo-lang.mojo.rpc.longrunning.OperationOrBuilder getValsOrBuilder(
      int index) {
    return vals_.get(index);
  }

  private byte memoizedIsInitialized = -1;
  @java.lang.Override
  public final boolean isInitialized() {
    byte isInitialized = memoizedIsInitialized;
    if (isInitialized == 1) return true;
    if (isInitialized == 0) return false;

    memoizedIsInitialized = 1;
    return true;
  }

  @java.lang.Override
  public void writeTo(com.google.protobuf.CodedOutputStream output)
                      throws java.io.IOException {
    for (int i = 0; i < vals_.size(); i++) {
      output.writeMessage(1, vals_.get(i));
    }
    unknownFields.writeTo(output);
  }

  @java.lang.Override
  public int getSerializedSize() {
    int size = memoizedSize;
    if (size != -1) return size;

    size = 0;
    for (int i = 0; i < vals_.size(); i++) {
      size += com.google.protobuf.CodedOutputStream
        .computeMessageSize(1, vals_.get(i));
    }
    size += unknownFields.getSerializedSize();
    memoizedSize = size;
    return size;
  }

  @java.lang.Override
  public boolean equals(final java.lang.Object obj) {
    if (obj == this) {
     return true;
    }
    if (!(obj instanceof org.mojo-lang.mojo.rpc.longrunning.Operations)) {
      return super.equals(obj);
    }
    org.mojo-lang.mojo.rpc.longrunning.Operations other = (org.mojo-lang.mojo.rpc.longrunning.Operations) obj;

    if (!getValsList()
        .equals(other.getValsList())) return false;
    if (!unknownFields.equals(other.unknownFields)) return false;
    return true;
  }

  @java.lang.Override
  public int hashCode() {
    if (memoizedHashCode != 0) {
      return memoizedHashCode;
    }
    int hash = 41;
    hash = (19 * hash) + getDescriptor().hashCode();
    if (getValsCount() > 0) {
      hash = (37 * hash) + VALS_FIELD_NUMBER;
      hash = (53 * hash) + getValsList().hashCode();
    }
    hash = (29 * hash) + unknownFields.hashCode();
    memoizedHashCode = hash;
    return hash;
  }

  public static org.mojo-lang.mojo.rpc.longrunning.Operations parseFrom(
      java.nio.ByteBuffer data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.mojo-lang.mojo.rpc.longrunning.Operations parseFrom(
      java.nio.ByteBuffer data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.mojo-lang.mojo.rpc.longrunning.Operations parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.mojo-lang.mojo.rpc.longrunning.Operations parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.mojo-lang.mojo.rpc.longrunning.Operations parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.mojo-lang.mojo.rpc.longrunning.Operations parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.mojo-lang.mojo.rpc.longrunning.Operations parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static org.mojo-lang.mojo.rpc.longrunning.Operations parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }
  public static org.mojo-lang.mojo.rpc.longrunning.Operations parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input);
  }
  public static org.mojo-lang.mojo.rpc.longrunning.Operations parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
  }
  public static org.mojo-lang.mojo.rpc.longrunning.Operations parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static org.mojo-lang.mojo.rpc.longrunning.Operations parseFrom(
      com.google.protobuf.CodedInputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }

  @java.lang.Override
  public Builder newBuilderForType() { return newBuilder(); }
  public static Builder newBuilder() {
    return DEFAULT_INSTANCE.toBuilder();
  }
  public static Builder newBuilder(org.mojo-lang.mojo.rpc.longrunning.Operations prototype) {
    return DEFAULT_INSTANCE.toBuilder().mergeFrom(prototype);
  }
  @java.lang.Override
  public Builder toBuilder() {
    return this == DEFAULT_INSTANCE
        ? new Builder() : new Builder().mergeFrom(this);
  }

  @java.lang.Override
  protected Builder newBuilderForType(
      com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
    Builder builder = new Builder(parent);
    return builder;
  }
  /**
   * Protobuf type {@code mojo.rpc.longrunning.Operations}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
      // @@protoc_insertion_point(builder_implements:mojo.rpc.longrunning.Operations)
      org.mojo-lang.mojo.rpc.longrunning.OperationsOrBuilder {
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return org.mojo-lang.mojo.rpc.longrunning.OperationServiceProto.internal_static_mojo_rpc_longrunning_Operations_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return org.mojo-lang.mojo.rpc.longrunning.OperationServiceProto.internal_static_mojo_rpc_longrunning_Operations_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              org.mojo-lang.mojo.rpc.longrunning.Operations.class, org.mojo-lang.mojo.rpc.longrunning.Operations.Builder.class);
    }

    // Construct using org.mojo-lang.mojo.rpc.longrunning.Operations.newBuilder()
    private Builder() {
      maybeForceBuilderInitialization();
    }

    private Builder(
        com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
      super(parent);
      maybeForceBuilderInitialization();
    }
    private void maybeForceBuilderInitialization() {
      if (com.google.protobuf.GeneratedMessageV3
              .alwaysUseFieldBuilders) {
        getValsFieldBuilder();
      }
    }
    @java.lang.Override
    public Builder clear() {
      super.clear();
      if (valsBuilder_ == null) {
        vals_ = java.util.Collections.emptyList();
        bitField0_ = (bitField0_ & ~0x00000001);
      } else {
        valsBuilder_.clear();
      }
      return this;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.Descriptor
        getDescriptorForType() {
      return org.mojo-lang.mojo.rpc.longrunning.OperationServiceProto.internal_static_mojo_rpc_longrunning_Operations_descriptor;
    }

    @java.lang.Override
    public org.mojo-lang.mojo.rpc.longrunning.Operations getDefaultInstanceForType() {
      return org.mojo-lang.mojo.rpc.longrunning.Operations.getDefaultInstance();
    }

    @java.lang.Override
    public org.mojo-lang.mojo.rpc.longrunning.Operations build() {
      org.mojo-lang.mojo.rpc.longrunning.Operations result = buildPartial();
      if (!result.isInitialized()) {
        throw newUninitializedMessageException(result);
      }
      return result;
    }

    @java.lang.Override
    public org.mojo-lang.mojo.rpc.longrunning.Operations buildPartial() {
      org.mojo-lang.mojo.rpc.longrunning.Operations result = new org.mojo-lang.mojo.rpc.longrunning.Operations(this);
      int from_bitField0_ = bitField0_;
      if (valsBuilder_ == null) {
        if (((bitField0_ & 0x00000001) != 0)) {
          vals_ = java.util.Collections.unmodifiableList(vals_);
          bitField0_ = (bitField0_ & ~0x00000001);
        }
        result.vals_ = vals_;
      } else {
        result.vals_ = valsBuilder_.build();
      }
      onBuilt();
      return result;
    }

    @java.lang.Override
    public Builder clone() {
      return super.clone();
    }
    @java.lang.Override
    public Builder setField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        java.lang.Object value) {
      return super.setField(field, value);
    }
    @java.lang.Override
    public Builder clearField(
        com.google.protobuf.Descriptors.FieldDescriptor field) {
      return super.clearField(field);
    }
    @java.lang.Override
    public Builder clearOneof(
        com.google.protobuf.Descriptors.OneofDescriptor oneof) {
      return super.clearOneof(oneof);
    }
    @java.lang.Override
    public Builder setRepeatedField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        int index, java.lang.Object value) {
      return super.setRepeatedField(field, index, value);
    }
    @java.lang.Override
    public Builder addRepeatedField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        java.lang.Object value) {
      return super.addRepeatedField(field, value);
    }
    @java.lang.Override
    public Builder mergeFrom(com.google.protobuf.Message other) {
      if (other instanceof org.mojo-lang.mojo.rpc.longrunning.Operations) {
        return mergeFrom((org.mojo-lang.mojo.rpc.longrunning.Operations)other);
      } else {
        super.mergeFrom(other);
        return this;
      }
    }

    public Builder mergeFrom(org.mojo-lang.mojo.rpc.longrunning.Operations other) {
      if (other == org.mojo-lang.mojo.rpc.longrunning.Operations.getDefaultInstance()) return this;
      if (valsBuilder_ == null) {
        if (!other.vals_.isEmpty()) {
          if (vals_.isEmpty()) {
            vals_ = other.vals_;
            bitField0_ = (bitField0_ & ~0x00000001);
          } else {
            ensureValsIsMutable();
            vals_.addAll(other.vals_);
          }
          onChanged();
        }
      } else {
        if (!other.vals_.isEmpty()) {
          if (valsBuilder_.isEmpty()) {
            valsBuilder_.dispose();
            valsBuilder_ = null;
            vals_ = other.vals_;
            bitField0_ = (bitField0_ & ~0x00000001);
            valsBuilder_ = 
              com.google.protobuf.GeneratedMessageV3.alwaysUseFieldBuilders ?
                 getValsFieldBuilder() : null;
          } else {
            valsBuilder_.addAllMessages(other.vals_);
          }
        }
      }
      this.mergeUnknownFields(other.unknownFields);
      onChanged();
      return this;
    }

    @java.lang.Override
    public final boolean isInitialized() {
      return true;
    }

    @java.lang.Override
    public Builder mergeFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      org.mojo-lang.mojo.rpc.longrunning.Operations parsedMessage = null;
      try {
        parsedMessage = PARSER.parsePartialFrom(input, extensionRegistry);
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        parsedMessage = (org.mojo-lang.mojo.rpc.longrunning.Operations) e.getUnfinishedMessage();
        throw e.unwrapIOException();
      } finally {
        if (parsedMessage != null) {
          mergeFrom(parsedMessage);
        }
      }
      return this;
    }
    private int bitField0_;

    private java.util.List<org.mojo-lang.mojo.rpc.longrunning.Operation> vals_ =
      java.util.Collections.emptyList();
    private void ensureValsIsMutable() {
      if (!((bitField0_ & 0x00000001) != 0)) {
        vals_ = new java.util.ArrayList<org.mojo-lang.mojo.rpc.longrunning.Operation>(vals_);
        bitField0_ |= 0x00000001;
       }
    }

    private com.google.protobuf.RepeatedFieldBuilderV3<
        org.mojo-lang.mojo.rpc.longrunning.Operation, org.mojo-lang.mojo.rpc.longrunning.Operation.Builder, org.mojo-lang.mojo.rpc.longrunning.OperationOrBuilder> valsBuilder_;

    /**
     * <code>repeated .mojo.rpc.longrunning.Operation vals = 1;</code>
     */
    public java.util.List<org.mojo-lang.mojo.rpc.longrunning.Operation> getValsList() {
      if (valsBuilder_ == null) {
        return java.util.Collections.unmodifiableList(vals_);
      } else {
        return valsBuilder_.getMessageList();
      }
    }
    /**
     * <code>repeated .mojo.rpc.longrunning.Operation vals = 1;</code>
     */
    public int getValsCount() {
      if (valsBuilder_ == null) {
        return vals_.size();
      } else {
        return valsBuilder_.getCount();
      }
    }
    /**
     * <code>repeated .mojo.rpc.longrunning.Operation vals = 1;</code>
     */
    public org.mojo-lang.mojo.rpc.longrunning.Operation getVals(int index) {
      if (valsBuilder_ == null) {
        return vals_.get(index);
      } else {
        return valsBuilder_.getMessage(index);
      }
    }
    /**
     * <code>repeated .mojo.rpc.longrunning.Operation vals = 1;</code>
     */
    public Builder setVals(
        int index, org.mojo-lang.mojo.rpc.longrunning.Operation value) {
      if (valsBuilder_ == null) {
        if (value == null) {
          throw new NullPointerException();
        }
        ensureValsIsMutable();
        vals_.set(index, value);
        onChanged();
      } else {
        valsBuilder_.setMessage(index, value);
      }
      return this;
    }
    /**
     * <code>repeated .mojo.rpc.longrunning.Operation vals = 1;</code>
     */
    public Builder setVals(
        int index, org.mojo-lang.mojo.rpc.longrunning.Operation.Builder builderForValue) {
      if (valsBuilder_ == null) {
        ensureValsIsMutable();
        vals_.set(index, builderForValue.build());
        onChanged();
      } else {
        valsBuilder_.setMessage(index, builderForValue.build());
      }
      return this;
    }
    /**
     * <code>repeated .mojo.rpc.longrunning.Operation vals = 1;</code>
     */
    public Builder addVals(org.mojo-lang.mojo.rpc.longrunning.Operation value) {
      if (valsBuilder_ == null) {
        if (value == null) {
          throw new NullPointerException();
        }
        ensureValsIsMutable();
        vals_.add(value);
        onChanged();
      } else {
        valsBuilder_.addMessage(value);
      }
      return this;
    }
    /**
     * <code>repeated .mojo.rpc.longrunning.Operation vals = 1;</code>
     */
    public Builder addVals(
        int index, org.mojo-lang.mojo.rpc.longrunning.Operation value) {
      if (valsBuilder_ == null) {
        if (value == null) {
          throw new NullPointerException();
        }
        ensureValsIsMutable();
        vals_.add(index, value);
        onChanged();
      } else {
        valsBuilder_.addMessage(index, value);
      }
      return this;
    }
    /**
     * <code>repeated .mojo.rpc.longrunning.Operation vals = 1;</code>
     */
    public Builder addVals(
        org.mojo-lang.mojo.rpc.longrunning.Operation.Builder builderForValue) {
      if (valsBuilder_ == null) {
        ensureValsIsMutable();
        vals_.add(builderForValue.build());
        onChanged();
      } else {
        valsBuilder_.addMessage(builderForValue.build());
      }
      return this;
    }
    /**
     * <code>repeated .mojo.rpc.longrunning.Operation vals = 1;</code>
     */
    public Builder addVals(
        int index, org.mojo-lang.mojo.rpc.longrunning.Operation.Builder builderForValue) {
      if (valsBuilder_ == null) {
        ensureValsIsMutable();
        vals_.add(index, builderForValue.build());
        onChanged();
      } else {
        valsBuilder_.addMessage(index, builderForValue.build());
      }
      return this;
    }
    /**
     * <code>repeated .mojo.rpc.longrunning.Operation vals = 1;</code>
     */
    public Builder addAllVals(
        java.lang.Iterable<? extends org.mojo-lang.mojo.rpc.longrunning.Operation> values) {
      if (valsBuilder_ == null) {
        ensureValsIsMutable();
        com.google.protobuf.AbstractMessageLite.Builder.addAll(
            values, vals_);
        onChanged();
      } else {
        valsBuilder_.addAllMessages(values);
      }
      return this;
    }
    /**
     * <code>repeated .mojo.rpc.longrunning.Operation vals = 1;</code>
     */
    public Builder clearVals() {
      if (valsBuilder_ == null) {
        vals_ = java.util.Collections.emptyList();
        bitField0_ = (bitField0_ & ~0x00000001);
        onChanged();
      } else {
        valsBuilder_.clear();
      }
      return this;
    }
    /**
     * <code>repeated .mojo.rpc.longrunning.Operation vals = 1;</code>
     */
    public Builder removeVals(int index) {
      if (valsBuilder_ == null) {
        ensureValsIsMutable();
        vals_.remove(index);
        onChanged();
      } else {
        valsBuilder_.remove(index);
      }
      return this;
    }
    /**
     * <code>repeated .mojo.rpc.longrunning.Operation vals = 1;</code>
     */
    public org.mojo-lang.mojo.rpc.longrunning.Operation.Builder getValsBuilder(
        int index) {
      return getValsFieldBuilder().getBuilder(index);
    }
    /**
     * <code>repeated .mojo.rpc.longrunning.Operation vals = 1;</code>
     */
    public org.mojo-lang.mojo.rpc.longrunning.OperationOrBuilder getValsOrBuilder(
        int index) {
      if (valsBuilder_ == null) {
        return vals_.get(index);  } else {
        return valsBuilder_.getMessageOrBuilder(index);
      }
    }
    /**
     * <code>repeated .mojo.rpc.longrunning.Operation vals = 1;</code>
     */
    public java.util.List<? extends org.mojo-lang.mojo.rpc.longrunning.OperationOrBuilder> 
         getValsOrBuilderList() {
      if (valsBuilder_ != null) {
        return valsBuilder_.getMessageOrBuilderList();
      } else {
        return java.util.Collections.unmodifiableList(vals_);
      }
    }
    /**
     * <code>repeated .mojo.rpc.longrunning.Operation vals = 1;</code>
     */
    public org.mojo-lang.mojo.rpc.longrunning.Operation.Builder addValsBuilder() {
      return getValsFieldBuilder().addBuilder(
          org.mojo-lang.mojo.rpc.longrunning.Operation.getDefaultInstance());
    }
    /**
     * <code>repeated .mojo.rpc.longrunning.Operation vals = 1;</code>
     */
    public org.mojo-lang.mojo.rpc.longrunning.Operation.Builder addValsBuilder(
        int index) {
      return getValsFieldBuilder().addBuilder(
          index, org.mojo-lang.mojo.rpc.longrunning.Operation.getDefaultInstance());
    }
    /**
     * <code>repeated .mojo.rpc.longrunning.Operation vals = 1;</code>
     */
    public java.util.List<org.mojo-lang.mojo.rpc.longrunning.Operation.Builder> 
         getValsBuilderList() {
      return getValsFieldBuilder().getBuilderList();
    }
    private com.google.protobuf.RepeatedFieldBuilderV3<
        org.mojo-lang.mojo.rpc.longrunning.Operation, org.mojo-lang.mojo.rpc.longrunning.Operation.Builder, org.mojo-lang.mojo.rpc.longrunning.OperationOrBuilder> 
        getValsFieldBuilder() {
      if (valsBuilder_ == null) {
        valsBuilder_ = new com.google.protobuf.RepeatedFieldBuilderV3<
            org.mojo-lang.mojo.rpc.longrunning.Operation, org.mojo-lang.mojo.rpc.longrunning.Operation.Builder, org.mojo-lang.mojo.rpc.longrunning.OperationOrBuilder>(
                vals_,
                ((bitField0_ & 0x00000001) != 0),
                getParentForChildren(),
                isClean());
        vals_ = null;
      }
      return valsBuilder_;
    }
    @java.lang.Override
    public final Builder setUnknownFields(
        final com.google.protobuf.UnknownFieldSet unknownFields) {
      return super.setUnknownFields(unknownFields);
    }

    @java.lang.Override
    public final Builder mergeUnknownFields(
        final com.google.protobuf.UnknownFieldSet unknownFields) {
      return super.mergeUnknownFields(unknownFields);
    }


    // @@protoc_insertion_point(builder_scope:mojo.rpc.longrunning.Operations)
  }

  // @@protoc_insertion_point(class_scope:mojo.rpc.longrunning.Operations)
  private static final org.mojo-lang.mojo.rpc.longrunning.Operations DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new org.mojo-lang.mojo.rpc.longrunning.Operations();
  }

  public static org.mojo-lang.mojo.rpc.longrunning.Operations getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static final com.google.protobuf.Parser<Operations>
      PARSER = new com.google.protobuf.AbstractParser<Operations>() {
    @java.lang.Override
    public Operations parsePartialFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return new Operations(input, extensionRegistry);
    }
  };

  public static com.google.protobuf.Parser<Operations> parser() {
    return PARSER;
  }

  @java.lang.Override
  public com.google.protobuf.Parser<Operations> getParserForType() {
    return PARSER;
  }

  @java.lang.Override
  public org.mojo-lang.mojo.rpc.longrunning.Operations getDefaultInstanceForType() {
    return DEFAULT_INSTANCE;
  }

}

