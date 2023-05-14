///
//  Generated code. Do not modify.
//  source: backend/grpc/service.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:async' as $async;
import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

class PriceRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'PriceRequest', createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'ticker')
    ..hasRequiredFields = false
  ;

  PriceRequest._() : super();
  factory PriceRequest({
    $core.String? ticker,
  }) {
    final _result = create();
    if (ticker != null) {
      _result.ticker = ticker;
    }
    return _result;
  }
  factory PriceRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory PriceRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  PriceRequest clone() => PriceRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  PriceRequest copyWith(void Function(PriceRequest) updates) => super.copyWith((message) => updates(message as PriceRequest)) as PriceRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static PriceRequest create() => PriceRequest._();
  PriceRequest createEmptyInstance() => create();
  static $pb.PbList<PriceRequest> createRepeated() => $pb.PbList<PriceRequest>();
  @$core.pragma('dart2js:noInline')
  static PriceRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<PriceRequest>(create);
  static PriceRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get ticker => $_getSZ(0);
  @$pb.TagNumber(1)
  set ticker($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasTicker() => $_has(0);
  @$pb.TagNumber(1)
  void clearTicker() => clearField(1);
}

class PriceResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'PriceResponse', createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'ticker')
    ..a<$core.double>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'price', $pb.PbFieldType.OF)
    ..hasRequiredFields = false
  ;

  PriceResponse._() : super();
  factory PriceResponse({
    $core.String? ticker,
    $core.double? price,
  }) {
    final _result = create();
    if (ticker != null) {
      _result.ticker = ticker;
    }
    if (price != null) {
      _result.price = price;
    }
    return _result;
  }
  factory PriceResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory PriceResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  PriceResponse clone() => PriceResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  PriceResponse copyWith(void Function(PriceResponse) updates) => super.copyWith((message) => updates(message as PriceResponse)) as PriceResponse; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static PriceResponse create() => PriceResponse._();
  PriceResponse createEmptyInstance() => create();
  static $pb.PbList<PriceResponse> createRepeated() => $pb.PbList<PriceResponse>();
  @$core.pragma('dart2js:noInline')
  static PriceResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<PriceResponse>(create);
  static PriceResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get ticker => $_getSZ(0);
  @$pb.TagNumber(1)
  set ticker($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasTicker() => $_has(0);
  @$pb.TagNumber(1)
  void clearTicker() => clearField(1);

  @$pb.TagNumber(2)
  $core.double get price => $_getN(1);
  @$pb.TagNumber(2)
  set price($core.double v) { $_setFloat(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasPrice() => $_has(1);
  @$pb.TagNumber(2)
  void clearPrice() => clearField(2);
}

class PriceFetcherApi {
  $pb.RpcClient _client;
  PriceFetcherApi(this._client);

  $async.Future<PriceResponse> fetchPrice($pb.ClientContext? ctx, PriceRequest request) {
    var emptyResponse = PriceResponse();
    return _client.invoke<PriceResponse>(ctx, 'PriceFetcher', 'FetchPrice', request, emptyResponse);
  }
}

