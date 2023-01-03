/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { BetInfo } from "./bet_info";
import { LotteryInfo } from "./lottery_info";
import { Params } from "./params";
import { StoredBet } from "./stored_bet";
import { StoredLottery } from "./stored_lottery";

export const protobufPackage = "naruto0913.lottery.lottery";

/** GenesisState defines the lottery module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  lotteryInfo: LotteryInfo | undefined;
  storedLotteryList: StoredLottery[];
  storedBetList: StoredBet[];
  /** this line is used by starport scaffolding # genesis/proto/state */
  betInfo: BetInfo | undefined;
}

function createBaseGenesisState(): GenesisState {
  return { params: undefined, lotteryInfo: undefined, storedLotteryList: [], storedBetList: [], betInfo: undefined };
}

export const GenesisState = {
  encode(message: GenesisState, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    if (message.lotteryInfo !== undefined) {
      LotteryInfo.encode(message.lotteryInfo, writer.uint32(18).fork()).ldelim();
    }
    for (const v of message.storedLotteryList) {
      StoredLottery.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    for (const v of message.storedBetList) {
      StoredBet.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    if (message.betInfo !== undefined) {
      BetInfo.encode(message.betInfo, writer.uint32(42).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGenesisState();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 2:
          message.lotteryInfo = LotteryInfo.decode(reader, reader.uint32());
          break;
        case 3:
          message.storedLotteryList.push(StoredLottery.decode(reader, reader.uint32()));
          break;
        case 4:
          message.storedBetList.push(StoredBet.decode(reader, reader.uint32()));
          break;
        case 5:
          message.betInfo = BetInfo.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    return {
      params: isSet(object.params) ? Params.fromJSON(object.params) : undefined,
      lotteryInfo: isSet(object.lotteryInfo) ? LotteryInfo.fromJSON(object.lotteryInfo) : undefined,
      storedLotteryList: Array.isArray(object?.storedLotteryList)
        ? object.storedLotteryList.map((e: any) => StoredLottery.fromJSON(e))
        : [],
      storedBetList: Array.isArray(object?.storedBetList)
        ? object.storedBetList.map((e: any) => StoredBet.fromJSON(e))
        : [],
      betInfo: isSet(object.betInfo) ? BetInfo.fromJSON(object.betInfo) : undefined,
    };
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined && (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    message.lotteryInfo !== undefined
      && (obj.lotteryInfo = message.lotteryInfo ? LotteryInfo.toJSON(message.lotteryInfo) : undefined);
    if (message.storedLotteryList) {
      obj.storedLotteryList = message.storedLotteryList.map((e) => e ? StoredLottery.toJSON(e) : undefined);
    } else {
      obj.storedLotteryList = [];
    }
    if (message.storedBetList) {
      obj.storedBetList = message.storedBetList.map((e) => e ? StoredBet.toJSON(e) : undefined);
    } else {
      obj.storedBetList = [];
    }
    message.betInfo !== undefined && (obj.betInfo = message.betInfo ? BetInfo.toJSON(message.betInfo) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GenesisState>, I>>(object: I): GenesisState {
    const message = createBaseGenesisState();
    message.params = (object.params !== undefined && object.params !== null)
      ? Params.fromPartial(object.params)
      : undefined;
    message.lotteryInfo = (object.lotteryInfo !== undefined && object.lotteryInfo !== null)
      ? LotteryInfo.fromPartial(object.lotteryInfo)
      : undefined;
    message.storedLotteryList = object.storedLotteryList?.map((e) => StoredLottery.fromPartial(e)) || [];
    message.storedBetList = object.storedBetList?.map((e) => StoredBet.fromPartial(e)) || [];
    message.betInfo = (object.betInfo !== undefined && object.betInfo !== null)
      ? BetInfo.fromPartial(object.betInfo)
      : undefined;
    return message;
  },
};

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
