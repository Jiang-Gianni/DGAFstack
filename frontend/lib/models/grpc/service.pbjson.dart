///
//  Generated code. Do not modify.
//  source: backend/grpc/service.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
@$core.Deprecated('Use priceRequestDescriptor instead')
const PriceRequest$json = const {
  '1': 'PriceRequest',
  '2': const [
    const {'1': 'ticker', '3': 1, '4': 1, '5': 9, '10': 'ticker'},
  ],
};

/// Descriptor for `PriceRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List priceRequestDescriptor = $convert.base64Decode('CgxQcmljZVJlcXVlc3QSFgoGdGlja2VyGAEgASgJUgZ0aWNrZXI=');
@$core.Deprecated('Use priceResponseDescriptor instead')
const PriceResponse$json = const {
  '1': 'PriceResponse',
  '2': const [
    const {'1': 'ticker', '3': 1, '4': 1, '5': 9, '10': 'ticker'},
    const {'1': 'price', '3': 2, '4': 1, '5': 2, '10': 'price'},
  ],
};

/// Descriptor for `PriceResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List priceResponseDescriptor = $convert.base64Decode('Cg1QcmljZVJlc3BvbnNlEhYKBnRpY2tlchgBIAEoCVIGdGlja2VyEhQKBXByaWNlGAIgASgCUgVwcmljZQ==');
const $core.Map<$core.String, $core.dynamic> PriceFetcherServiceBase$json = const {
  '1': 'PriceFetcher',
  '2': const [
    const {'1': 'FetchPrice', '2': '.PriceRequest', '3': '.PriceResponse'},
  ],
};

@$core.Deprecated('Use priceFetcherServiceDescriptor instead')
const $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> PriceFetcherServiceBase$messageJson = const {
  '.PriceRequest': PriceRequest$json,
  '.PriceResponse': PriceResponse$json,
};

/// Descriptor for `PriceFetcher`. Decode as a `google.protobuf.ServiceDescriptorProto`.
final $typed_data.Uint8List priceFetcherServiceDescriptor = $convert.base64Decode('CgxQcmljZUZldGNoZXISKwoKRmV0Y2hQcmljZRINLlByaWNlUmVxdWVzdBoOLlByaWNlUmVzcG9uc2U=');
