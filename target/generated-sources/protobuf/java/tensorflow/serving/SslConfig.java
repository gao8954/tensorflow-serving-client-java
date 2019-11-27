// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: tensorflow_serving/config/ssl_config.proto

package tensorflow.serving;

public final class SslConfig {
  private SslConfig() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  public interface SSLConfigOrBuilder extends
      // @@protoc_insertion_point(interface_extends:tensorflow.serving.SSLConfig)
      com.google.protobuf.MessageOrBuilder {

    /**
     * <pre>
     * private server key for SSL
     * </pre>
     *
     * <code>string server_key = 1;</code>
     */
    java.lang.String getServerKey();
    /**
     * <pre>
     * private server key for SSL
     * </pre>
     *
     * <code>string server_key = 1;</code>
     */
    com.google.protobuf.ByteString
        getServerKeyBytes();

    /**
     * <pre>
     * public server certificate
     * </pre>
     *
     * <code>string server_cert = 2;</code>
     */
    java.lang.String getServerCert();
    /**
     * <pre>
     * public server certificate
     * </pre>
     *
     * <code>string server_cert = 2;</code>
     */
    com.google.protobuf.ByteString
        getServerCertBytes();

    /**
     * <pre>
     *  custom certificate authority
     * </pre>
     *
     * <code>string custom_ca = 3;</code>
     */
    java.lang.String getCustomCa();
    /**
     * <pre>
     *  custom certificate authority
     * </pre>
     *
     * <code>string custom_ca = 3;</code>
     */
    com.google.protobuf.ByteString
        getCustomCaBytes();

    /**
     * <pre>
     * valid client certificate required ?
     * </pre>
     *
     * <code>bool client_verify = 4;</code>
     */
    boolean getClientVerify();
  }
  /**
   * <pre>
   * Configuration for a secure gRPC channel
   * </pre>
   *
   * Protobuf type {@code tensorflow.serving.SSLConfig}
   */
  public  static final class SSLConfig extends
      com.google.protobuf.GeneratedMessageV3 implements
      // @@protoc_insertion_point(message_implements:tensorflow.serving.SSLConfig)
      SSLConfigOrBuilder {
  private static final long serialVersionUID = 0L;
    // Use SSLConfig.newBuilder() to construct.
    private SSLConfig(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
      super(builder);
    }
    private SSLConfig() {
      serverKey_ = "";
      serverCert_ = "";
      customCa_ = "";
      clientVerify_ = false;
    }

    @java.lang.Override
    public final com.google.protobuf.UnknownFieldSet
    getUnknownFields() {
      return this.unknownFields;
    }
    private SSLConfig(
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
            default: {
              if (!parseUnknownFieldProto3(
                  input, unknownFields, extensionRegistry, tag)) {
                done = true;
              }
              break;
            }
            case 10: {
              java.lang.String s = input.readStringRequireUtf8();

              serverKey_ = s;
              break;
            }
            case 18: {
              java.lang.String s = input.readStringRequireUtf8();

              serverCert_ = s;
              break;
            }
            case 26: {
              java.lang.String s = input.readStringRequireUtf8();

              customCa_ = s;
              break;
            }
            case 32: {

              clientVerify_ = input.readBool();
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
        this.unknownFields = unknownFields.build();
        makeExtensionsImmutable();
      }
    }
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return tensorflow.serving.SslConfig.internal_static_tensorflow_serving_SSLConfig_descriptor;
    }

    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return tensorflow.serving.SslConfig.internal_static_tensorflow_serving_SSLConfig_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              tensorflow.serving.SslConfig.SSLConfig.class, tensorflow.serving.SslConfig.SSLConfig.Builder.class);
    }

    public static final int SERVER_KEY_FIELD_NUMBER = 1;
    private volatile java.lang.Object serverKey_;
    /**
     * <pre>
     * private server key for SSL
     * </pre>
     *
     * <code>string server_key = 1;</code>
     */
    public java.lang.String getServerKey() {
      java.lang.Object ref = serverKey_;
      if (ref instanceof java.lang.String) {
        return (java.lang.String) ref;
      } else {
        com.google.protobuf.ByteString bs = 
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        serverKey_ = s;
        return s;
      }
    }
    /**
     * <pre>
     * private server key for SSL
     * </pre>
     *
     * <code>string server_key = 1;</code>
     */
    public com.google.protobuf.ByteString
        getServerKeyBytes() {
      java.lang.Object ref = serverKey_;
      if (ref instanceof java.lang.String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        serverKey_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }

    public static final int SERVER_CERT_FIELD_NUMBER = 2;
    private volatile java.lang.Object serverCert_;
    /**
     * <pre>
     * public server certificate
     * </pre>
     *
     * <code>string server_cert = 2;</code>
     */
    public java.lang.String getServerCert() {
      java.lang.Object ref = serverCert_;
      if (ref instanceof java.lang.String) {
        return (java.lang.String) ref;
      } else {
        com.google.protobuf.ByteString bs = 
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        serverCert_ = s;
        return s;
      }
    }
    /**
     * <pre>
     * public server certificate
     * </pre>
     *
     * <code>string server_cert = 2;</code>
     */
    public com.google.protobuf.ByteString
        getServerCertBytes() {
      java.lang.Object ref = serverCert_;
      if (ref instanceof java.lang.String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        serverCert_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }

    public static final int CUSTOM_CA_FIELD_NUMBER = 3;
    private volatile java.lang.Object customCa_;
    /**
     * <pre>
     *  custom certificate authority
     * </pre>
     *
     * <code>string custom_ca = 3;</code>
     */
    public java.lang.String getCustomCa() {
      java.lang.Object ref = customCa_;
      if (ref instanceof java.lang.String) {
        return (java.lang.String) ref;
      } else {
        com.google.protobuf.ByteString bs = 
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        customCa_ = s;
        return s;
      }
    }
    /**
     * <pre>
     *  custom certificate authority
     * </pre>
     *
     * <code>string custom_ca = 3;</code>
     */
    public com.google.protobuf.ByteString
        getCustomCaBytes() {
      java.lang.Object ref = customCa_;
      if (ref instanceof java.lang.String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        customCa_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }

    public static final int CLIENT_VERIFY_FIELD_NUMBER = 4;
    private boolean clientVerify_;
    /**
     * <pre>
     * valid client certificate required ?
     * </pre>
     *
     * <code>bool client_verify = 4;</code>
     */
    public boolean getClientVerify() {
      return clientVerify_;
    }

    private byte memoizedIsInitialized = -1;
    public final boolean isInitialized() {
      byte isInitialized = memoizedIsInitialized;
      if (isInitialized == 1) return true;
      if (isInitialized == 0) return false;

      memoizedIsInitialized = 1;
      return true;
    }

    public void writeTo(com.google.protobuf.CodedOutputStream output)
                        throws java.io.IOException {
      if (!getServerKeyBytes().isEmpty()) {
        com.google.protobuf.GeneratedMessageV3.writeString(output, 1, serverKey_);
      }
      if (!getServerCertBytes().isEmpty()) {
        com.google.protobuf.GeneratedMessageV3.writeString(output, 2, serverCert_);
      }
      if (!getCustomCaBytes().isEmpty()) {
        com.google.protobuf.GeneratedMessageV3.writeString(output, 3, customCa_);
      }
      if (clientVerify_ != false) {
        output.writeBool(4, clientVerify_);
      }
      unknownFields.writeTo(output);
    }

    public int getSerializedSize() {
      int size = memoizedSize;
      if (size != -1) return size;

      size = 0;
      if (!getServerKeyBytes().isEmpty()) {
        size += com.google.protobuf.GeneratedMessageV3.computeStringSize(1, serverKey_);
      }
      if (!getServerCertBytes().isEmpty()) {
        size += com.google.protobuf.GeneratedMessageV3.computeStringSize(2, serverCert_);
      }
      if (!getCustomCaBytes().isEmpty()) {
        size += com.google.protobuf.GeneratedMessageV3.computeStringSize(3, customCa_);
      }
      if (clientVerify_ != false) {
        size += com.google.protobuf.CodedOutputStream
          .computeBoolSize(4, clientVerify_);
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
      if (!(obj instanceof tensorflow.serving.SslConfig.SSLConfig)) {
        return super.equals(obj);
      }
      tensorflow.serving.SslConfig.SSLConfig other = (tensorflow.serving.SslConfig.SSLConfig) obj;

      boolean result = true;
      result = result && getServerKey()
          .equals(other.getServerKey());
      result = result && getServerCert()
          .equals(other.getServerCert());
      result = result && getCustomCa()
          .equals(other.getCustomCa());
      result = result && (getClientVerify()
          == other.getClientVerify());
      result = result && unknownFields.equals(other.unknownFields);
      return result;
    }

    @java.lang.Override
    public int hashCode() {
      if (memoizedHashCode != 0) {
        return memoizedHashCode;
      }
      int hash = 41;
      hash = (19 * hash) + getDescriptor().hashCode();
      hash = (37 * hash) + SERVER_KEY_FIELD_NUMBER;
      hash = (53 * hash) + getServerKey().hashCode();
      hash = (37 * hash) + SERVER_CERT_FIELD_NUMBER;
      hash = (53 * hash) + getServerCert().hashCode();
      hash = (37 * hash) + CUSTOM_CA_FIELD_NUMBER;
      hash = (53 * hash) + getCustomCa().hashCode();
      hash = (37 * hash) + CLIENT_VERIFY_FIELD_NUMBER;
      hash = (53 * hash) + com.google.protobuf.Internal.hashBoolean(
          getClientVerify());
      hash = (29 * hash) + unknownFields.hashCode();
      memoizedHashCode = hash;
      return hash;
    }

    public static tensorflow.serving.SslConfig.SSLConfig parseFrom(
        java.nio.ByteBuffer data)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data);
    }
    public static tensorflow.serving.SslConfig.SSLConfig parseFrom(
        java.nio.ByteBuffer data,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data, extensionRegistry);
    }
    public static tensorflow.serving.SslConfig.SSLConfig parseFrom(
        com.google.protobuf.ByteString data)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data);
    }
    public static tensorflow.serving.SslConfig.SSLConfig parseFrom(
        com.google.protobuf.ByteString data,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data, extensionRegistry);
    }
    public static tensorflow.serving.SslConfig.SSLConfig parseFrom(byte[] data)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data);
    }
    public static tensorflow.serving.SslConfig.SSLConfig parseFrom(
        byte[] data,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data, extensionRegistry);
    }
    public static tensorflow.serving.SslConfig.SSLConfig parseFrom(java.io.InputStream input)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseWithIOException(PARSER, input);
    }
    public static tensorflow.serving.SslConfig.SSLConfig parseFrom(
        java.io.InputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseWithIOException(PARSER, input, extensionRegistry);
    }
    public static tensorflow.serving.SslConfig.SSLConfig parseDelimitedFrom(java.io.InputStream input)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseDelimitedWithIOException(PARSER, input);
    }
    public static tensorflow.serving.SslConfig.SSLConfig parseDelimitedFrom(
        java.io.InputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
    }
    public static tensorflow.serving.SslConfig.SSLConfig parseFrom(
        com.google.protobuf.CodedInputStream input)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseWithIOException(PARSER, input);
    }
    public static tensorflow.serving.SslConfig.SSLConfig parseFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseWithIOException(PARSER, input, extensionRegistry);
    }

    public Builder newBuilderForType() { return newBuilder(); }
    public static Builder newBuilder() {
      return DEFAULT_INSTANCE.toBuilder();
    }
    public static Builder newBuilder(tensorflow.serving.SslConfig.SSLConfig prototype) {
      return DEFAULT_INSTANCE.toBuilder().mergeFrom(prototype);
    }
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
     * <pre>
     * Configuration for a secure gRPC channel
     * </pre>
     *
     * Protobuf type {@code tensorflow.serving.SSLConfig}
     */
    public static final class Builder extends
        com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
        // @@protoc_insertion_point(builder_implements:tensorflow.serving.SSLConfig)
        tensorflow.serving.SslConfig.SSLConfigOrBuilder {
      public static final com.google.protobuf.Descriptors.Descriptor
          getDescriptor() {
        return tensorflow.serving.SslConfig.internal_static_tensorflow_serving_SSLConfig_descriptor;
      }

      protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
          internalGetFieldAccessorTable() {
        return tensorflow.serving.SslConfig.internal_static_tensorflow_serving_SSLConfig_fieldAccessorTable
            .ensureFieldAccessorsInitialized(
                tensorflow.serving.SslConfig.SSLConfig.class, tensorflow.serving.SslConfig.SSLConfig.Builder.class);
      }

      // Construct using tensorflow.serving.SslConfig.SSLConfig.newBuilder()
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
        }
      }
      public Builder clear() {
        super.clear();
        serverKey_ = "";

        serverCert_ = "";

        customCa_ = "";

        clientVerify_ = false;

        return this;
      }

      public com.google.protobuf.Descriptors.Descriptor
          getDescriptorForType() {
        return tensorflow.serving.SslConfig.internal_static_tensorflow_serving_SSLConfig_descriptor;
      }

      public tensorflow.serving.SslConfig.SSLConfig getDefaultInstanceForType() {
        return tensorflow.serving.SslConfig.SSLConfig.getDefaultInstance();
      }

      public tensorflow.serving.SslConfig.SSLConfig build() {
        tensorflow.serving.SslConfig.SSLConfig result = buildPartial();
        if (!result.isInitialized()) {
          throw newUninitializedMessageException(result);
        }
        return result;
      }

      public tensorflow.serving.SslConfig.SSLConfig buildPartial() {
        tensorflow.serving.SslConfig.SSLConfig result = new tensorflow.serving.SslConfig.SSLConfig(this);
        result.serverKey_ = serverKey_;
        result.serverCert_ = serverCert_;
        result.customCa_ = customCa_;
        result.clientVerify_ = clientVerify_;
        onBuilt();
        return result;
      }

      public Builder clone() {
        return (Builder) super.clone();
      }
      public Builder setField(
          com.google.protobuf.Descriptors.FieldDescriptor field,
          java.lang.Object value) {
        return (Builder) super.setField(field, value);
      }
      public Builder clearField(
          com.google.protobuf.Descriptors.FieldDescriptor field) {
        return (Builder) super.clearField(field);
      }
      public Builder clearOneof(
          com.google.protobuf.Descriptors.OneofDescriptor oneof) {
        return (Builder) super.clearOneof(oneof);
      }
      public Builder setRepeatedField(
          com.google.protobuf.Descriptors.FieldDescriptor field,
          int index, java.lang.Object value) {
        return (Builder) super.setRepeatedField(field, index, value);
      }
      public Builder addRepeatedField(
          com.google.protobuf.Descriptors.FieldDescriptor field,
          java.lang.Object value) {
        return (Builder) super.addRepeatedField(field, value);
      }
      public Builder mergeFrom(com.google.protobuf.Message other) {
        if (other instanceof tensorflow.serving.SslConfig.SSLConfig) {
          return mergeFrom((tensorflow.serving.SslConfig.SSLConfig)other);
        } else {
          super.mergeFrom(other);
          return this;
        }
      }

      public Builder mergeFrom(tensorflow.serving.SslConfig.SSLConfig other) {
        if (other == tensorflow.serving.SslConfig.SSLConfig.getDefaultInstance()) return this;
        if (!other.getServerKey().isEmpty()) {
          serverKey_ = other.serverKey_;
          onChanged();
        }
        if (!other.getServerCert().isEmpty()) {
          serverCert_ = other.serverCert_;
          onChanged();
        }
        if (!other.getCustomCa().isEmpty()) {
          customCa_ = other.customCa_;
          onChanged();
        }
        if (other.getClientVerify() != false) {
          setClientVerify(other.getClientVerify());
        }
        this.mergeUnknownFields(other.unknownFields);
        onChanged();
        return this;
      }

      public final boolean isInitialized() {
        return true;
      }

      public Builder mergeFrom(
          com.google.protobuf.CodedInputStream input,
          com.google.protobuf.ExtensionRegistryLite extensionRegistry)
          throws java.io.IOException {
        tensorflow.serving.SslConfig.SSLConfig parsedMessage = null;
        try {
          parsedMessage = PARSER.parsePartialFrom(input, extensionRegistry);
        } catch (com.google.protobuf.InvalidProtocolBufferException e) {
          parsedMessage = (tensorflow.serving.SslConfig.SSLConfig) e.getUnfinishedMessage();
          throw e.unwrapIOException();
        } finally {
          if (parsedMessage != null) {
            mergeFrom(parsedMessage);
          }
        }
        return this;
      }

      private java.lang.Object serverKey_ = "";
      /**
       * <pre>
       * private server key for SSL
       * </pre>
       *
       * <code>string server_key = 1;</code>
       */
      public java.lang.String getServerKey() {
        java.lang.Object ref = serverKey_;
        if (!(ref instanceof java.lang.String)) {
          com.google.protobuf.ByteString bs =
              (com.google.protobuf.ByteString) ref;
          java.lang.String s = bs.toStringUtf8();
          serverKey_ = s;
          return s;
        } else {
          return (java.lang.String) ref;
        }
      }
      /**
       * <pre>
       * private server key for SSL
       * </pre>
       *
       * <code>string server_key = 1;</code>
       */
      public com.google.protobuf.ByteString
          getServerKeyBytes() {
        java.lang.Object ref = serverKey_;
        if (ref instanceof String) {
          com.google.protobuf.ByteString b = 
              com.google.protobuf.ByteString.copyFromUtf8(
                  (java.lang.String) ref);
          serverKey_ = b;
          return b;
        } else {
          return (com.google.protobuf.ByteString) ref;
        }
      }
      /**
       * <pre>
       * private server key for SSL
       * </pre>
       *
       * <code>string server_key = 1;</code>
       */
      public Builder setServerKey(
          java.lang.String value) {
        if (value == null) {
    throw new NullPointerException();
  }
  
        serverKey_ = value;
        onChanged();
        return this;
      }
      /**
       * <pre>
       * private server key for SSL
       * </pre>
       *
       * <code>string server_key = 1;</code>
       */
      public Builder clearServerKey() {
        
        serverKey_ = getDefaultInstance().getServerKey();
        onChanged();
        return this;
      }
      /**
       * <pre>
       * private server key for SSL
       * </pre>
       *
       * <code>string server_key = 1;</code>
       */
      public Builder setServerKeyBytes(
          com.google.protobuf.ByteString value) {
        if (value == null) {
    throw new NullPointerException();
  }
  checkByteStringIsUtf8(value);
        
        serverKey_ = value;
        onChanged();
        return this;
      }

      private java.lang.Object serverCert_ = "";
      /**
       * <pre>
       * public server certificate
       * </pre>
       *
       * <code>string server_cert = 2;</code>
       */
      public java.lang.String getServerCert() {
        java.lang.Object ref = serverCert_;
        if (!(ref instanceof java.lang.String)) {
          com.google.protobuf.ByteString bs =
              (com.google.protobuf.ByteString) ref;
          java.lang.String s = bs.toStringUtf8();
          serverCert_ = s;
          return s;
        } else {
          return (java.lang.String) ref;
        }
      }
      /**
       * <pre>
       * public server certificate
       * </pre>
       *
       * <code>string server_cert = 2;</code>
       */
      public com.google.protobuf.ByteString
          getServerCertBytes() {
        java.lang.Object ref = serverCert_;
        if (ref instanceof String) {
          com.google.protobuf.ByteString b = 
              com.google.protobuf.ByteString.copyFromUtf8(
                  (java.lang.String) ref);
          serverCert_ = b;
          return b;
        } else {
          return (com.google.protobuf.ByteString) ref;
        }
      }
      /**
       * <pre>
       * public server certificate
       * </pre>
       *
       * <code>string server_cert = 2;</code>
       */
      public Builder setServerCert(
          java.lang.String value) {
        if (value == null) {
    throw new NullPointerException();
  }
  
        serverCert_ = value;
        onChanged();
        return this;
      }
      /**
       * <pre>
       * public server certificate
       * </pre>
       *
       * <code>string server_cert = 2;</code>
       */
      public Builder clearServerCert() {
        
        serverCert_ = getDefaultInstance().getServerCert();
        onChanged();
        return this;
      }
      /**
       * <pre>
       * public server certificate
       * </pre>
       *
       * <code>string server_cert = 2;</code>
       */
      public Builder setServerCertBytes(
          com.google.protobuf.ByteString value) {
        if (value == null) {
    throw new NullPointerException();
  }
  checkByteStringIsUtf8(value);
        
        serverCert_ = value;
        onChanged();
        return this;
      }

      private java.lang.Object customCa_ = "";
      /**
       * <pre>
       *  custom certificate authority
       * </pre>
       *
       * <code>string custom_ca = 3;</code>
       */
      public java.lang.String getCustomCa() {
        java.lang.Object ref = customCa_;
        if (!(ref instanceof java.lang.String)) {
          com.google.protobuf.ByteString bs =
              (com.google.protobuf.ByteString) ref;
          java.lang.String s = bs.toStringUtf8();
          customCa_ = s;
          return s;
        } else {
          return (java.lang.String) ref;
        }
      }
      /**
       * <pre>
       *  custom certificate authority
       * </pre>
       *
       * <code>string custom_ca = 3;</code>
       */
      public com.google.protobuf.ByteString
          getCustomCaBytes() {
        java.lang.Object ref = customCa_;
        if (ref instanceof String) {
          com.google.protobuf.ByteString b = 
              com.google.protobuf.ByteString.copyFromUtf8(
                  (java.lang.String) ref);
          customCa_ = b;
          return b;
        } else {
          return (com.google.protobuf.ByteString) ref;
        }
      }
      /**
       * <pre>
       *  custom certificate authority
       * </pre>
       *
       * <code>string custom_ca = 3;</code>
       */
      public Builder setCustomCa(
          java.lang.String value) {
        if (value == null) {
    throw new NullPointerException();
  }
  
        customCa_ = value;
        onChanged();
        return this;
      }
      /**
       * <pre>
       *  custom certificate authority
       * </pre>
       *
       * <code>string custom_ca = 3;</code>
       */
      public Builder clearCustomCa() {
        
        customCa_ = getDefaultInstance().getCustomCa();
        onChanged();
        return this;
      }
      /**
       * <pre>
       *  custom certificate authority
       * </pre>
       *
       * <code>string custom_ca = 3;</code>
       */
      public Builder setCustomCaBytes(
          com.google.protobuf.ByteString value) {
        if (value == null) {
    throw new NullPointerException();
  }
  checkByteStringIsUtf8(value);
        
        customCa_ = value;
        onChanged();
        return this;
      }

      private boolean clientVerify_ ;
      /**
       * <pre>
       * valid client certificate required ?
       * </pre>
       *
       * <code>bool client_verify = 4;</code>
       */
      public boolean getClientVerify() {
        return clientVerify_;
      }
      /**
       * <pre>
       * valid client certificate required ?
       * </pre>
       *
       * <code>bool client_verify = 4;</code>
       */
      public Builder setClientVerify(boolean value) {
        
        clientVerify_ = value;
        onChanged();
        return this;
      }
      /**
       * <pre>
       * valid client certificate required ?
       * </pre>
       *
       * <code>bool client_verify = 4;</code>
       */
      public Builder clearClientVerify() {
        
        clientVerify_ = false;
        onChanged();
        return this;
      }
      public final Builder setUnknownFields(
          final com.google.protobuf.UnknownFieldSet unknownFields) {
        return super.setUnknownFieldsProto3(unknownFields);
      }

      public final Builder mergeUnknownFields(
          final com.google.protobuf.UnknownFieldSet unknownFields) {
        return super.mergeUnknownFields(unknownFields);
      }


      // @@protoc_insertion_point(builder_scope:tensorflow.serving.SSLConfig)
    }

    // @@protoc_insertion_point(class_scope:tensorflow.serving.SSLConfig)
    private static final tensorflow.serving.SslConfig.SSLConfig DEFAULT_INSTANCE;
    static {
      DEFAULT_INSTANCE = new tensorflow.serving.SslConfig.SSLConfig();
    }

    public static tensorflow.serving.SslConfig.SSLConfig getDefaultInstance() {
      return DEFAULT_INSTANCE;
    }

    private static final com.google.protobuf.Parser<SSLConfig>
        PARSER = new com.google.protobuf.AbstractParser<SSLConfig>() {
      public SSLConfig parsePartialFrom(
          com.google.protobuf.CodedInputStream input,
          com.google.protobuf.ExtensionRegistryLite extensionRegistry)
          throws com.google.protobuf.InvalidProtocolBufferException {
        return new SSLConfig(input, extensionRegistry);
      }
    };

    public static com.google.protobuf.Parser<SSLConfig> parser() {
      return PARSER;
    }

    @java.lang.Override
    public com.google.protobuf.Parser<SSLConfig> getParserForType() {
      return PARSER;
    }

    public tensorflow.serving.SslConfig.SSLConfig getDefaultInstanceForType() {
      return DEFAULT_INSTANCE;
    }

  }

  private static final com.google.protobuf.Descriptors.Descriptor
    internal_static_tensorflow_serving_SSLConfig_descriptor;
  private static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_tensorflow_serving_SSLConfig_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n*tensorflow_serving/config/ssl_config.p" +
      "roto\022\022tensorflow.serving\"^\n\tSSLConfig\022\022\n" +
      "\nserver_key\030\001 \001(\t\022\023\n\013server_cert\030\002 \001(\t\022\021" +
      "\n\tcustom_ca\030\003 \001(\t\022\025\n\rclient_verify\030\004 \001(\010" +
      "B\003\370\001\001b\006proto3"
    };
    com.google.protobuf.Descriptors.FileDescriptor.InternalDescriptorAssigner assigner =
        new com.google.protobuf.Descriptors.FileDescriptor.    InternalDescriptorAssigner() {
          public com.google.protobuf.ExtensionRegistry assignDescriptors(
              com.google.protobuf.Descriptors.FileDescriptor root) {
            descriptor = root;
            return null;
          }
        };
    com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
        }, assigner);
    internal_static_tensorflow_serving_SSLConfig_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_tensorflow_serving_SSLConfig_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_tensorflow_serving_SSLConfig_descriptor,
        new java.lang.String[] { "ServerKey", "ServerCert", "CustomCa", "ClientVerify", });
  }

  // @@protoc_insertion_point(outer_class_scope)
}
