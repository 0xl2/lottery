/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "naruto0913.lottery.lottery";

export interface StoredBet {
  index: string;
  userAddr: string;
  lotteryIndex: number;
  orderIndex: number;
  betAmount: number;
}

function createBaseStoredBet(): StoredBet {
  return { index: "", userAddr: "", lotteryIndex: 0, orderIndex: 0, betAmount: 0 };
}

export const StoredBet = {
  encode(message: StoredBet, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.userAddr !== "") {
      writer.uint32(18).string(message.userAddr);
    }
    if (message.lotteryIndex !== 0) {
      writer.uint32(24).uint64(message.lotteryIndex);
    }
    if (message.orderIndex !== 0) {
      writer.uint32(32).uint64(message.orderIndex);
    }
    if (message.betAmount !== 0) {
      writer.uint32(40).uint64(message.betAmount);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): StoredBet {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseStoredBet();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.userAddr = reader.string();
          break;
        case 3:
          message.lotteryIndex = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.orderIndex = longToNumber(reader.uint64() as Long);
          break;
        case 5:
          message.betAmount = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): StoredBet {
    return {
      index: isSet(object.index) ? String(object.index) : "",
      userAddr: isSet(object.userAddr) ? String(object.userAddr) : "",
      lotteryIndex: isSet(object.lotteryIndex) ? Number(object.lotteryIndex) : 0,
      orderIndex: isSet(object.orderIndex) ? Number(object.orderIndex) : 0,
      betAmount: isSet(object.betAmount) ? Number(object.betAmount) : 0,
    };
  },

  toJSON(message: StoredBet): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.userAddr !== undefined && (obj.userAddr = message.userAddr);
    message.lotteryIndex !== undefined && (obj.lotteryIndex = Math.round(message.lotteryIndex));
    message.orderIndex !== undefined && (obj.orderIndex = Math.round(message.orderIndex));
    message.betAmount !== undefined && (obj.betAmount = Math.round(message.betAmount));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<StoredBet>, I>>(object: I): StoredBet {
    const message = createBaseStoredBet();
    message.index = object.index ?? "";
    message.userAddr = object.userAddr ?? "";
    message.lotteryIndex = object.lotteryIndex ?? 0;
    message.orderIndex = object.orderIndex ?? 0;
    message.betAmount = object.betAmount ?? 0;
    return message;
  },
};

declare var self: any | undefined;
declare var window: any | undefined;
declare var global: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") {
    return globalThis;
  }
  if (typeof self !== "undefined") {
    return self;
  }
  if (typeof window !== "undefined") {
    return window;
  }
  if (typeof global !== "undefined") {
    return global;
  }
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
