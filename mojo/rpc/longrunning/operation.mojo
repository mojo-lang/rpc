// Copyright 2020 Google LLC
// [source](https://github.com/googleapis/googleapis/google/longrunning/operations.proto)
//
// Copyright 2021 Mojo-lang.org
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/// This resource represents a long-running operation that is the result of a
/// network API call.
type Operation {
  /// The server-assigned name, which is only unique within the same service that
  /// originally returns it. If you use the default HTTP mapping, the
  /// `name` should be a resource name ending with `operations/{unique_id}`.
  name: String @1 @key

  /// Service-specific metadata associated with the operation.  It typically
  /// contains progress information and common metadata such as create time.
  /// Some services might not provide such metadata.  Any method that returns a
  /// long-running operation should document the metadata type, if any.
  metadata: Any @2

  /// If the value is `false`, it means the operation is still in progress.
  /// If `true`, the operation is completed, and either `error` or `response` is
  /// available.
  done: Bool @3

  /// The operation result, which can be either an `error` or a valid `response`.
  /// If `done` == `false`, neither `error` nor `response` is set.
  /// If `done` == `true`, exactly one of `error` or `response` is set.
  /// The error result of the operation in case of failure or cancellation.
  error: Error @4

  /// The normal response of the operation in case of success.  If the original
  /// method returns no data on success, such as `Delete`, the response is
  /// `google.protobuf.Empty`.  If the original method is standard
  /// `Get`/`Create`/`Update`, the response should be the resource.  For other
  /// methods, the response should have the type `XxxResponse`, where `Xxx`
  /// is the original method name.  For example, if the original method name
  /// is `TakeSnapshot()`, the inferred response type is
  /// `TakeSnapshotResponse`.
  response: Any @5

  /// the create timestamp for the operation created.
  create_time: Timestamp @14

  /// the updated timestamp for the operation when update the progression information.
  update_time: Timestamp @15
}
