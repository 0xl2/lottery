/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { PageRequest, PageResponse } from "../../cosmos/base/query/v1beta1/pagination";
import { BetInfo } from "./bet_info";
import { LotteryInfo } from "./lottery_info";
import { Params } from "./params";
import { StoredBet } from "./stored_bet";
import { StoredLottery } from "./stored_lottery";

export const protobufPackage = "naruto0913.lottery.lottery";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {
}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetLotteryInfoRequest {
}

export interface QueryGetLotteryInfoResponse {
  LotteryInfo: LotteryInfo | undefined;
}

export interface QueryGetStoredLotteryRequest {
  index: string;
}

export interface QueryGetStoredLotteryResponse {
  storedLottery: StoredLottery | undefined;
}

export interface QueryAllStoredLotteryRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllStoredLotteryResponse {
  storedLottery: StoredLottery[];
  pagination: PageResponse | undefined;
}

export interface QueryGetStoredBetRequest {
  index: string;
}

export interface QueryGetStoredBetResponse {
  storedBet: StoredBet | undefined;
}

export interface QueryAllStoredBetRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllStoredBetResponse {
  storedBet: StoredBet[];
  pagination: PageResponse | undefined;
}

export interface QueryGetBetInfoRequest {
}

export interface QueryGetBetInfoResponse {
  BetInfo: BetInfo | undefined;
}

function createBaseQueryParamsRequest(): QueryParamsRequest {
  return {};
}

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryParamsRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): QueryParamsRequest {
    return {};
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryParamsRequest>, I>>(_: I): QueryParamsRequest {
    const message = createBaseQueryParamsRequest();
    return message;
  },
};

function createBaseQueryParamsResponse(): QueryParamsResponse {
  return { params: undefined };
}

export const QueryParamsResponse = {
  encode(message: QueryParamsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryParamsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    return { params: isSet(object.params) ? Params.fromJSON(object.params) : undefined };
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined && (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryParamsResponse>, I>>(object: I): QueryParamsResponse {
    const message = createBaseQueryParamsResponse();
    message.params = (object.params !== undefined && object.params !== null)
      ? Params.fromPartial(object.params)
      : undefined;
    return message;
  },
};

function createBaseQueryGetLotteryInfoRequest(): QueryGetLotteryInfoRequest {
  return {};
}

export const QueryGetLotteryInfoRequest = {
  encode(_: QueryGetLotteryInfoRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetLotteryInfoRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetLotteryInfoRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): QueryGetLotteryInfoRequest {
    return {};
  },

  toJSON(_: QueryGetLotteryInfoRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetLotteryInfoRequest>, I>>(_: I): QueryGetLotteryInfoRequest {
    const message = createBaseQueryGetLotteryInfoRequest();
    return message;
  },
};

function createBaseQueryGetLotteryInfoResponse(): QueryGetLotteryInfoResponse {
  return { LotteryInfo: undefined };
}

export const QueryGetLotteryInfoResponse = {
  encode(message: QueryGetLotteryInfoResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.LotteryInfo !== undefined) {
      LotteryInfo.encode(message.LotteryInfo, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetLotteryInfoResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetLotteryInfoResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.LotteryInfo = LotteryInfo.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetLotteryInfoResponse {
    return { LotteryInfo: isSet(object.LotteryInfo) ? LotteryInfo.fromJSON(object.LotteryInfo) : undefined };
  },

  toJSON(message: QueryGetLotteryInfoResponse): unknown {
    const obj: any = {};
    message.LotteryInfo !== undefined
      && (obj.LotteryInfo = message.LotteryInfo ? LotteryInfo.toJSON(message.LotteryInfo) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetLotteryInfoResponse>, I>>(object: I): QueryGetLotteryInfoResponse {
    const message = createBaseQueryGetLotteryInfoResponse();
    message.LotteryInfo = (object.LotteryInfo !== undefined && object.LotteryInfo !== null)
      ? LotteryInfo.fromPartial(object.LotteryInfo)
      : undefined;
    return message;
  },
};

function createBaseQueryGetStoredLotteryRequest(): QueryGetStoredLotteryRequest {
  return { index: "" };
}

export const QueryGetStoredLotteryRequest = {
  encode(message: QueryGetStoredLotteryRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetStoredLotteryRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetStoredLotteryRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetStoredLotteryRequest {
    return { index: isSet(object.index) ? String(object.index) : "" };
  },

  toJSON(message: QueryGetStoredLotteryRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetStoredLotteryRequest>, I>>(object: I): QueryGetStoredLotteryRequest {
    const message = createBaseQueryGetStoredLotteryRequest();
    message.index = object.index ?? "";
    return message;
  },
};

function createBaseQueryGetStoredLotteryResponse(): QueryGetStoredLotteryResponse {
  return { storedLottery: undefined };
}

export const QueryGetStoredLotteryResponse = {
  encode(message: QueryGetStoredLotteryResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.storedLottery !== undefined) {
      StoredLottery.encode(message.storedLottery, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetStoredLotteryResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetStoredLotteryResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.storedLottery = StoredLottery.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetStoredLotteryResponse {
    return { storedLottery: isSet(object.storedLottery) ? StoredLottery.fromJSON(object.storedLottery) : undefined };
  },

  toJSON(message: QueryGetStoredLotteryResponse): unknown {
    const obj: any = {};
    message.storedLottery !== undefined
      && (obj.storedLottery = message.storedLottery ? StoredLottery.toJSON(message.storedLottery) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetStoredLotteryResponse>, I>>(
    object: I,
  ): QueryGetStoredLotteryResponse {
    const message = createBaseQueryGetStoredLotteryResponse();
    message.storedLottery = (object.storedLottery !== undefined && object.storedLottery !== null)
      ? StoredLottery.fromPartial(object.storedLottery)
      : undefined;
    return message;
  },
};

function createBaseQueryAllStoredLotteryRequest(): QueryAllStoredLotteryRequest {
  return { pagination: undefined };
}

export const QueryAllStoredLotteryRequest = {
  encode(message: QueryAllStoredLotteryRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllStoredLotteryRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllStoredLotteryRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllStoredLotteryRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllStoredLotteryRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllStoredLotteryRequest>, I>>(object: I): QueryAllStoredLotteryRequest {
    const message = createBaseQueryAllStoredLotteryRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllStoredLotteryResponse(): QueryAllStoredLotteryResponse {
  return { storedLottery: [], pagination: undefined };
}

export const QueryAllStoredLotteryResponse = {
  encode(message: QueryAllStoredLotteryResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.storedLottery) {
      StoredLottery.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllStoredLotteryResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllStoredLotteryResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.storedLottery.push(StoredLottery.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllStoredLotteryResponse {
    return {
      storedLottery: Array.isArray(object?.storedLottery)
        ? object.storedLottery.map((e: any) => StoredLottery.fromJSON(e))
        : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllStoredLotteryResponse): unknown {
    const obj: any = {};
    if (message.storedLottery) {
      obj.storedLottery = message.storedLottery.map((e) => e ? StoredLottery.toJSON(e) : undefined);
    } else {
      obj.storedLottery = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllStoredLotteryResponse>, I>>(
    object: I,
  ): QueryAllStoredLotteryResponse {
    const message = createBaseQueryAllStoredLotteryResponse();
    message.storedLottery = object.storedLottery?.map((e) => StoredLottery.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryGetStoredBetRequest(): QueryGetStoredBetRequest {
  return { index: "" };
}

export const QueryGetStoredBetRequest = {
  encode(message: QueryGetStoredBetRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetStoredBetRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetStoredBetRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetStoredBetRequest {
    return { index: isSet(object.index) ? String(object.index) : "" };
  },

  toJSON(message: QueryGetStoredBetRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetStoredBetRequest>, I>>(object: I): QueryGetStoredBetRequest {
    const message = createBaseQueryGetStoredBetRequest();
    message.index = object.index ?? "";
    return message;
  },
};

function createBaseQueryGetStoredBetResponse(): QueryGetStoredBetResponse {
  return { storedBet: undefined };
}

export const QueryGetStoredBetResponse = {
  encode(message: QueryGetStoredBetResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.storedBet !== undefined) {
      StoredBet.encode(message.storedBet, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetStoredBetResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetStoredBetResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.storedBet = StoredBet.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetStoredBetResponse {
    return { storedBet: isSet(object.storedBet) ? StoredBet.fromJSON(object.storedBet) : undefined };
  },

  toJSON(message: QueryGetStoredBetResponse): unknown {
    const obj: any = {};
    message.storedBet !== undefined
      && (obj.storedBet = message.storedBet ? StoredBet.toJSON(message.storedBet) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetStoredBetResponse>, I>>(object: I): QueryGetStoredBetResponse {
    const message = createBaseQueryGetStoredBetResponse();
    message.storedBet = (object.storedBet !== undefined && object.storedBet !== null)
      ? StoredBet.fromPartial(object.storedBet)
      : undefined;
    return message;
  },
};

function createBaseQueryAllStoredBetRequest(): QueryAllStoredBetRequest {
  return { pagination: undefined };
}

export const QueryAllStoredBetRequest = {
  encode(message: QueryAllStoredBetRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllStoredBetRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllStoredBetRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllStoredBetRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllStoredBetRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllStoredBetRequest>, I>>(object: I): QueryAllStoredBetRequest {
    const message = createBaseQueryAllStoredBetRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllStoredBetResponse(): QueryAllStoredBetResponse {
  return { storedBet: [], pagination: undefined };
}

export const QueryAllStoredBetResponse = {
  encode(message: QueryAllStoredBetResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.storedBet) {
      StoredBet.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllStoredBetResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllStoredBetResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.storedBet.push(StoredBet.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllStoredBetResponse {
    return {
      storedBet: Array.isArray(object?.storedBet) ? object.storedBet.map((e: any) => StoredBet.fromJSON(e)) : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllStoredBetResponse): unknown {
    const obj: any = {};
    if (message.storedBet) {
      obj.storedBet = message.storedBet.map((e) => e ? StoredBet.toJSON(e) : undefined);
    } else {
      obj.storedBet = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllStoredBetResponse>, I>>(object: I): QueryAllStoredBetResponse {
    const message = createBaseQueryAllStoredBetResponse();
    message.storedBet = object.storedBet?.map((e) => StoredBet.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryGetBetInfoRequest(): QueryGetBetInfoRequest {
  return {};
}

export const QueryGetBetInfoRequest = {
  encode(_: QueryGetBetInfoRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetBetInfoRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetBetInfoRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): QueryGetBetInfoRequest {
    return {};
  },

  toJSON(_: QueryGetBetInfoRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetBetInfoRequest>, I>>(_: I): QueryGetBetInfoRequest {
    const message = createBaseQueryGetBetInfoRequest();
    return message;
  },
};

function createBaseQueryGetBetInfoResponse(): QueryGetBetInfoResponse {
  return { BetInfo: undefined };
}

export const QueryGetBetInfoResponse = {
  encode(message: QueryGetBetInfoResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.BetInfo !== undefined) {
      BetInfo.encode(message.BetInfo, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetBetInfoResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetBetInfoResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.BetInfo = BetInfo.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetBetInfoResponse {
    return { BetInfo: isSet(object.BetInfo) ? BetInfo.fromJSON(object.BetInfo) : undefined };
  },

  toJSON(message: QueryGetBetInfoResponse): unknown {
    const obj: any = {};
    message.BetInfo !== undefined && (obj.BetInfo = message.BetInfo ? BetInfo.toJSON(message.BetInfo) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetBetInfoResponse>, I>>(object: I): QueryGetBetInfoResponse {
    const message = createBaseQueryGetBetInfoResponse();
    message.BetInfo = (object.BetInfo !== undefined && object.BetInfo !== null)
      ? BetInfo.fromPartial(object.BetInfo)
      : undefined;
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a LotteryInfo by index. */
  LotteryInfo(request: QueryGetLotteryInfoRequest): Promise<QueryGetLotteryInfoResponse>;
  /** Queries a StoredLottery by index. */
  StoredLottery(request: QueryGetStoredLotteryRequest): Promise<QueryGetStoredLotteryResponse>;
  /** Queries a list of StoredLottery items. */
  StoredLotteryAll(request: QueryAllStoredLotteryRequest): Promise<QueryAllStoredLotteryResponse>;
  /** Queries a StoredBet by index. */
  StoredBet(request: QueryGetStoredBetRequest): Promise<QueryGetStoredBetResponse>;
  /** Queries a list of StoredBet items. */
  StoredBetAll(request: QueryAllStoredBetRequest): Promise<QueryAllStoredBetResponse>;
  /** Queries a BetInfo by index. */
  BetInfo(request: QueryGetBetInfoRequest): Promise<QueryGetBetInfoResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.Params = this.Params.bind(this);
    this.LotteryInfo = this.LotteryInfo.bind(this);
    this.StoredLottery = this.StoredLottery.bind(this);
    this.StoredLotteryAll = this.StoredLotteryAll.bind(this);
    this.StoredBet = this.StoredBet.bind(this);
    this.StoredBetAll = this.StoredBetAll.bind(this);
    this.BetInfo = this.BetInfo.bind(this);
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request("naruto0913.lottery.lottery.Query", "Params", data);
    return promise.then((data) => QueryParamsResponse.decode(new _m0.Reader(data)));
  }

  LotteryInfo(request: QueryGetLotteryInfoRequest): Promise<QueryGetLotteryInfoResponse> {
    const data = QueryGetLotteryInfoRequest.encode(request).finish();
    const promise = this.rpc.request("naruto0913.lottery.lottery.Query", "LotteryInfo", data);
    return promise.then((data) => QueryGetLotteryInfoResponse.decode(new _m0.Reader(data)));
  }

  StoredLottery(request: QueryGetStoredLotteryRequest): Promise<QueryGetStoredLotteryResponse> {
    const data = QueryGetStoredLotteryRequest.encode(request).finish();
    const promise = this.rpc.request("naruto0913.lottery.lottery.Query", "StoredLottery", data);
    return promise.then((data) => QueryGetStoredLotteryResponse.decode(new _m0.Reader(data)));
  }

  StoredLotteryAll(request: QueryAllStoredLotteryRequest): Promise<QueryAllStoredLotteryResponse> {
    const data = QueryAllStoredLotteryRequest.encode(request).finish();
    const promise = this.rpc.request("naruto0913.lottery.lottery.Query", "StoredLotteryAll", data);
    return promise.then((data) => QueryAllStoredLotteryResponse.decode(new _m0.Reader(data)));
  }

  StoredBet(request: QueryGetStoredBetRequest): Promise<QueryGetStoredBetResponse> {
    const data = QueryGetStoredBetRequest.encode(request).finish();
    const promise = this.rpc.request("naruto0913.lottery.lottery.Query", "StoredBet", data);
    return promise.then((data) => QueryGetStoredBetResponse.decode(new _m0.Reader(data)));
  }

  StoredBetAll(request: QueryAllStoredBetRequest): Promise<QueryAllStoredBetResponse> {
    const data = QueryAllStoredBetRequest.encode(request).finish();
    const promise = this.rpc.request("naruto0913.lottery.lottery.Query", "StoredBetAll", data);
    return promise.then((data) => QueryAllStoredBetResponse.decode(new _m0.Reader(data)));
  }

  BetInfo(request: QueryGetBetInfoRequest): Promise<QueryGetBetInfoResponse> {
    const data = QueryGetBetInfoRequest.encode(request).finish();
    const promise = this.rpc.request("naruto0913.lottery.lottery.Query", "BetInfo", data);
    return promise.then((data) => QueryGetBetInfoResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
