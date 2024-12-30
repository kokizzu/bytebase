// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.3.0
//   protoc               unknown
// source: store/task_run.proto

/* eslint-disable */
import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";
import Long from "long";
import { Timestamp } from "../google/protobuf/timestamp";
import { Position } from "./common";

export const protobufPackage = "bytebase.store";

export interface TaskRunResult {
  detail: string;
  /** Format: instances/{instance}/databases/{database}/changelogs/{changelog} */
  changelog: string;
  version: string;
  startPosition: TaskRunResult_Position | undefined;
  endPosition:
    | TaskRunResult_Position
    | undefined;
  /** The uid of the export archive. */
  exportArchiveUid: number;
  /** The prior backup detail that will be used to rollback the task run. */
  priorBackupDetail: PriorBackupDetail | undefined;
}

/** The following fields are used for error reporting. */
export interface TaskRunResult_Position {
  line: number;
  column: number;
}

export interface PriorBackupDetail {
  items: PriorBackupDetail_Item[];
}

export interface PriorBackupDetail_Item {
  /** The original table information. */
  sourceTable:
    | PriorBackupDetail_Item_Table
    | undefined;
  /** The target backup table information. */
  targetTable: PriorBackupDetail_Item_Table | undefined;
  startPosition: Position | undefined;
  endPosition: Position | undefined;
}

export interface PriorBackupDetail_Item_Table {
  /**
   * The database information.
   * Format: instances/{instance}/databases/{database}
   */
  database: string;
  schema: string;
  table: string;
}

export interface SchedulerInfo {
  reportTime: Timestamp | undefined;
  waitingCause: SchedulerInfo_WaitingCause | undefined;
}

export interface SchedulerInfo_WaitingCause {
  connectionLimit?: boolean | undefined;
  taskUid?: number | undefined;
}

function createBaseTaskRunResult(): TaskRunResult {
  return {
    detail: "",
    changelog: "",
    version: "",
    startPosition: undefined,
    endPosition: undefined,
    exportArchiveUid: 0,
    priorBackupDetail: undefined,
  };
}

export const TaskRunResult: MessageFns<TaskRunResult> = {
  encode(message: TaskRunResult, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.detail !== "") {
      writer.uint32(10).string(message.detail);
    }
    if (message.changelog !== "") {
      writer.uint32(66).string(message.changelog);
    }
    if (message.version !== "") {
      writer.uint32(26).string(message.version);
    }
    if (message.startPosition !== undefined) {
      TaskRunResult_Position.encode(message.startPosition, writer.uint32(34).fork()).join();
    }
    if (message.endPosition !== undefined) {
      TaskRunResult_Position.encode(message.endPosition, writer.uint32(42).fork()).join();
    }
    if (message.exportArchiveUid !== 0) {
      writer.uint32(48).int32(message.exportArchiveUid);
    }
    if (message.priorBackupDetail !== undefined) {
      PriorBackupDetail.encode(message.priorBackupDetail, writer.uint32(58).fork()).join();
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): TaskRunResult {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseTaskRunResult();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.detail = reader.string();
          continue;
        }
        case 8: {
          if (tag !== 66) {
            break;
          }

          message.changelog = reader.string();
          continue;
        }
        case 3: {
          if (tag !== 26) {
            break;
          }

          message.version = reader.string();
          continue;
        }
        case 4: {
          if (tag !== 34) {
            break;
          }

          message.startPosition = TaskRunResult_Position.decode(reader, reader.uint32());
          continue;
        }
        case 5: {
          if (tag !== 42) {
            break;
          }

          message.endPosition = TaskRunResult_Position.decode(reader, reader.uint32());
          continue;
        }
        case 6: {
          if (tag !== 48) {
            break;
          }

          message.exportArchiveUid = reader.int32();
          continue;
        }
        case 7: {
          if (tag !== 58) {
            break;
          }

          message.priorBackupDetail = PriorBackupDetail.decode(reader, reader.uint32());
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): TaskRunResult {
    return {
      detail: isSet(object.detail) ? globalThis.String(object.detail) : "",
      changelog: isSet(object.changelog) ? globalThis.String(object.changelog) : "",
      version: isSet(object.version) ? globalThis.String(object.version) : "",
      startPosition: isSet(object.startPosition) ? TaskRunResult_Position.fromJSON(object.startPosition) : undefined,
      endPosition: isSet(object.endPosition) ? TaskRunResult_Position.fromJSON(object.endPosition) : undefined,
      exportArchiveUid: isSet(object.exportArchiveUid) ? globalThis.Number(object.exportArchiveUid) : 0,
      priorBackupDetail: isSet(object.priorBackupDetail)
        ? PriorBackupDetail.fromJSON(object.priorBackupDetail)
        : undefined,
    };
  },

  toJSON(message: TaskRunResult): unknown {
    const obj: any = {};
    if (message.detail !== "") {
      obj.detail = message.detail;
    }
    if (message.changelog !== "") {
      obj.changelog = message.changelog;
    }
    if (message.version !== "") {
      obj.version = message.version;
    }
    if (message.startPosition !== undefined) {
      obj.startPosition = TaskRunResult_Position.toJSON(message.startPosition);
    }
    if (message.endPosition !== undefined) {
      obj.endPosition = TaskRunResult_Position.toJSON(message.endPosition);
    }
    if (message.exportArchiveUid !== 0) {
      obj.exportArchiveUid = Math.round(message.exportArchiveUid);
    }
    if (message.priorBackupDetail !== undefined) {
      obj.priorBackupDetail = PriorBackupDetail.toJSON(message.priorBackupDetail);
    }
    return obj;
  },

  create(base?: DeepPartial<TaskRunResult>): TaskRunResult {
    return TaskRunResult.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<TaskRunResult>): TaskRunResult {
    const message = createBaseTaskRunResult();
    message.detail = object.detail ?? "";
    message.changelog = object.changelog ?? "";
    message.version = object.version ?? "";
    message.startPosition = (object.startPosition !== undefined && object.startPosition !== null)
      ? TaskRunResult_Position.fromPartial(object.startPosition)
      : undefined;
    message.endPosition = (object.endPosition !== undefined && object.endPosition !== null)
      ? TaskRunResult_Position.fromPartial(object.endPosition)
      : undefined;
    message.exportArchiveUid = object.exportArchiveUid ?? 0;
    message.priorBackupDetail = (object.priorBackupDetail !== undefined && object.priorBackupDetail !== null)
      ? PriorBackupDetail.fromPartial(object.priorBackupDetail)
      : undefined;
    return message;
  },
};

function createBaseTaskRunResult_Position(): TaskRunResult_Position {
  return { line: 0, column: 0 };
}

export const TaskRunResult_Position: MessageFns<TaskRunResult_Position> = {
  encode(message: TaskRunResult_Position, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.line !== 0) {
      writer.uint32(8).int32(message.line);
    }
    if (message.column !== 0) {
      writer.uint32(16).int32(message.column);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): TaskRunResult_Position {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseTaskRunResult_Position();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 8) {
            break;
          }

          message.line = reader.int32();
          continue;
        }
        case 2: {
          if (tag !== 16) {
            break;
          }

          message.column = reader.int32();
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): TaskRunResult_Position {
    return {
      line: isSet(object.line) ? globalThis.Number(object.line) : 0,
      column: isSet(object.column) ? globalThis.Number(object.column) : 0,
    };
  },

  toJSON(message: TaskRunResult_Position): unknown {
    const obj: any = {};
    if (message.line !== 0) {
      obj.line = Math.round(message.line);
    }
    if (message.column !== 0) {
      obj.column = Math.round(message.column);
    }
    return obj;
  },

  create(base?: DeepPartial<TaskRunResult_Position>): TaskRunResult_Position {
    return TaskRunResult_Position.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<TaskRunResult_Position>): TaskRunResult_Position {
    const message = createBaseTaskRunResult_Position();
    message.line = object.line ?? 0;
    message.column = object.column ?? 0;
    return message;
  },
};

function createBasePriorBackupDetail(): PriorBackupDetail {
  return { items: [] };
}

export const PriorBackupDetail: MessageFns<PriorBackupDetail> = {
  encode(message: PriorBackupDetail, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    for (const v of message.items) {
      PriorBackupDetail_Item.encode(v!, writer.uint32(10).fork()).join();
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): PriorBackupDetail {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePriorBackupDetail();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.items.push(PriorBackupDetail_Item.decode(reader, reader.uint32()));
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): PriorBackupDetail {
    return {
      items: globalThis.Array.isArray(object?.items)
        ? object.items.map((e: any) => PriorBackupDetail_Item.fromJSON(e))
        : [],
    };
  },

  toJSON(message: PriorBackupDetail): unknown {
    const obj: any = {};
    if (message.items?.length) {
      obj.items = message.items.map((e) => PriorBackupDetail_Item.toJSON(e));
    }
    return obj;
  },

  create(base?: DeepPartial<PriorBackupDetail>): PriorBackupDetail {
    return PriorBackupDetail.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<PriorBackupDetail>): PriorBackupDetail {
    const message = createBasePriorBackupDetail();
    message.items = object.items?.map((e) => PriorBackupDetail_Item.fromPartial(e)) || [];
    return message;
  },
};

function createBasePriorBackupDetail_Item(): PriorBackupDetail_Item {
  return { sourceTable: undefined, targetTable: undefined, startPosition: undefined, endPosition: undefined };
}

export const PriorBackupDetail_Item: MessageFns<PriorBackupDetail_Item> = {
  encode(message: PriorBackupDetail_Item, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.sourceTable !== undefined) {
      PriorBackupDetail_Item_Table.encode(message.sourceTable, writer.uint32(10).fork()).join();
    }
    if (message.targetTable !== undefined) {
      PriorBackupDetail_Item_Table.encode(message.targetTable, writer.uint32(18).fork()).join();
    }
    if (message.startPosition !== undefined) {
      Position.encode(message.startPosition, writer.uint32(26).fork()).join();
    }
    if (message.endPosition !== undefined) {
      Position.encode(message.endPosition, writer.uint32(34).fork()).join();
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): PriorBackupDetail_Item {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePriorBackupDetail_Item();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.sourceTable = PriorBackupDetail_Item_Table.decode(reader, reader.uint32());
          continue;
        }
        case 2: {
          if (tag !== 18) {
            break;
          }

          message.targetTable = PriorBackupDetail_Item_Table.decode(reader, reader.uint32());
          continue;
        }
        case 3: {
          if (tag !== 26) {
            break;
          }

          message.startPosition = Position.decode(reader, reader.uint32());
          continue;
        }
        case 4: {
          if (tag !== 34) {
            break;
          }

          message.endPosition = Position.decode(reader, reader.uint32());
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): PriorBackupDetail_Item {
    return {
      sourceTable: isSet(object.sourceTable) ? PriorBackupDetail_Item_Table.fromJSON(object.sourceTable) : undefined,
      targetTable: isSet(object.targetTable) ? PriorBackupDetail_Item_Table.fromJSON(object.targetTable) : undefined,
      startPosition: isSet(object.startPosition) ? Position.fromJSON(object.startPosition) : undefined,
      endPosition: isSet(object.endPosition) ? Position.fromJSON(object.endPosition) : undefined,
    };
  },

  toJSON(message: PriorBackupDetail_Item): unknown {
    const obj: any = {};
    if (message.sourceTable !== undefined) {
      obj.sourceTable = PriorBackupDetail_Item_Table.toJSON(message.sourceTable);
    }
    if (message.targetTable !== undefined) {
      obj.targetTable = PriorBackupDetail_Item_Table.toJSON(message.targetTable);
    }
    if (message.startPosition !== undefined) {
      obj.startPosition = Position.toJSON(message.startPosition);
    }
    if (message.endPosition !== undefined) {
      obj.endPosition = Position.toJSON(message.endPosition);
    }
    return obj;
  },

  create(base?: DeepPartial<PriorBackupDetail_Item>): PriorBackupDetail_Item {
    return PriorBackupDetail_Item.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<PriorBackupDetail_Item>): PriorBackupDetail_Item {
    const message = createBasePriorBackupDetail_Item();
    message.sourceTable = (object.sourceTable !== undefined && object.sourceTable !== null)
      ? PriorBackupDetail_Item_Table.fromPartial(object.sourceTable)
      : undefined;
    message.targetTable = (object.targetTable !== undefined && object.targetTable !== null)
      ? PriorBackupDetail_Item_Table.fromPartial(object.targetTable)
      : undefined;
    message.startPosition = (object.startPosition !== undefined && object.startPosition !== null)
      ? Position.fromPartial(object.startPosition)
      : undefined;
    message.endPosition = (object.endPosition !== undefined && object.endPosition !== null)
      ? Position.fromPartial(object.endPosition)
      : undefined;
    return message;
  },
};

function createBasePriorBackupDetail_Item_Table(): PriorBackupDetail_Item_Table {
  return { database: "", schema: "", table: "" };
}

export const PriorBackupDetail_Item_Table: MessageFns<PriorBackupDetail_Item_Table> = {
  encode(message: PriorBackupDetail_Item_Table, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.database !== "") {
      writer.uint32(10).string(message.database);
    }
    if (message.schema !== "") {
      writer.uint32(18).string(message.schema);
    }
    if (message.table !== "") {
      writer.uint32(26).string(message.table);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): PriorBackupDetail_Item_Table {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePriorBackupDetail_Item_Table();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.database = reader.string();
          continue;
        }
        case 2: {
          if (tag !== 18) {
            break;
          }

          message.schema = reader.string();
          continue;
        }
        case 3: {
          if (tag !== 26) {
            break;
          }

          message.table = reader.string();
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): PriorBackupDetail_Item_Table {
    return {
      database: isSet(object.database) ? globalThis.String(object.database) : "",
      schema: isSet(object.schema) ? globalThis.String(object.schema) : "",
      table: isSet(object.table) ? globalThis.String(object.table) : "",
    };
  },

  toJSON(message: PriorBackupDetail_Item_Table): unknown {
    const obj: any = {};
    if (message.database !== "") {
      obj.database = message.database;
    }
    if (message.schema !== "") {
      obj.schema = message.schema;
    }
    if (message.table !== "") {
      obj.table = message.table;
    }
    return obj;
  },

  create(base?: DeepPartial<PriorBackupDetail_Item_Table>): PriorBackupDetail_Item_Table {
    return PriorBackupDetail_Item_Table.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<PriorBackupDetail_Item_Table>): PriorBackupDetail_Item_Table {
    const message = createBasePriorBackupDetail_Item_Table();
    message.database = object.database ?? "";
    message.schema = object.schema ?? "";
    message.table = object.table ?? "";
    return message;
  },
};

function createBaseSchedulerInfo(): SchedulerInfo {
  return { reportTime: undefined, waitingCause: undefined };
}

export const SchedulerInfo: MessageFns<SchedulerInfo> = {
  encode(message: SchedulerInfo, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.reportTime !== undefined) {
      Timestamp.encode(message.reportTime, writer.uint32(10).fork()).join();
    }
    if (message.waitingCause !== undefined) {
      SchedulerInfo_WaitingCause.encode(message.waitingCause, writer.uint32(18).fork()).join();
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): SchedulerInfo {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSchedulerInfo();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.reportTime = Timestamp.decode(reader, reader.uint32());
          continue;
        }
        case 2: {
          if (tag !== 18) {
            break;
          }

          message.waitingCause = SchedulerInfo_WaitingCause.decode(reader, reader.uint32());
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): SchedulerInfo {
    return {
      reportTime: isSet(object.reportTime) ? fromJsonTimestamp(object.reportTime) : undefined,
      waitingCause: isSet(object.waitingCause) ? SchedulerInfo_WaitingCause.fromJSON(object.waitingCause) : undefined,
    };
  },

  toJSON(message: SchedulerInfo): unknown {
    const obj: any = {};
    if (message.reportTime !== undefined) {
      obj.reportTime = fromTimestamp(message.reportTime).toISOString();
    }
    if (message.waitingCause !== undefined) {
      obj.waitingCause = SchedulerInfo_WaitingCause.toJSON(message.waitingCause);
    }
    return obj;
  },

  create(base?: DeepPartial<SchedulerInfo>): SchedulerInfo {
    return SchedulerInfo.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<SchedulerInfo>): SchedulerInfo {
    const message = createBaseSchedulerInfo();
    message.reportTime = (object.reportTime !== undefined && object.reportTime !== null)
      ? Timestamp.fromPartial(object.reportTime)
      : undefined;
    message.waitingCause = (object.waitingCause !== undefined && object.waitingCause !== null)
      ? SchedulerInfo_WaitingCause.fromPartial(object.waitingCause)
      : undefined;
    return message;
  },
};

function createBaseSchedulerInfo_WaitingCause(): SchedulerInfo_WaitingCause {
  return { connectionLimit: undefined, taskUid: undefined };
}

export const SchedulerInfo_WaitingCause: MessageFns<SchedulerInfo_WaitingCause> = {
  encode(message: SchedulerInfo_WaitingCause, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.connectionLimit !== undefined) {
      writer.uint32(8).bool(message.connectionLimit);
    }
    if (message.taskUid !== undefined) {
      writer.uint32(16).int32(message.taskUid);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): SchedulerInfo_WaitingCause {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSchedulerInfo_WaitingCause();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 8) {
            break;
          }

          message.connectionLimit = reader.bool();
          continue;
        }
        case 2: {
          if (tag !== 16) {
            break;
          }

          message.taskUid = reader.int32();
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): SchedulerInfo_WaitingCause {
    return {
      connectionLimit: isSet(object.connectionLimit) ? globalThis.Boolean(object.connectionLimit) : undefined,
      taskUid: isSet(object.taskUid) ? globalThis.Number(object.taskUid) : undefined,
    };
  },

  toJSON(message: SchedulerInfo_WaitingCause): unknown {
    const obj: any = {};
    if (message.connectionLimit !== undefined) {
      obj.connectionLimit = message.connectionLimit;
    }
    if (message.taskUid !== undefined) {
      obj.taskUid = Math.round(message.taskUid);
    }
    return obj;
  },

  create(base?: DeepPartial<SchedulerInfo_WaitingCause>): SchedulerInfo_WaitingCause {
    return SchedulerInfo_WaitingCause.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<SchedulerInfo_WaitingCause>): SchedulerInfo_WaitingCause {
    const message = createBaseSchedulerInfo_WaitingCause();
    message.connectionLimit = object.connectionLimit ?? undefined;
    message.taskUid = object.taskUid ?? undefined;
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Long ? string | number | Long : T extends globalThis.Array<infer U> ? globalThis.Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

function toTimestamp(date: Date): Timestamp {
  const seconds = numberToLong(Math.trunc(date.getTime() / 1_000));
  const nanos = (date.getTime() % 1_000) * 1_000_000;
  return { seconds, nanos };
}

function fromTimestamp(t: Timestamp): Date {
  let millis = (t.seconds.toNumber() || 0) * 1_000;
  millis += (t.nanos || 0) / 1_000_000;
  return new globalThis.Date(millis);
}

function fromJsonTimestamp(o: any): Timestamp {
  if (o instanceof globalThis.Date) {
    return toTimestamp(o);
  } else if (typeof o === "string") {
    return toTimestamp(new globalThis.Date(o));
  } else {
    return Timestamp.fromJSON(o);
  }
}

function numberToLong(number: number) {
  return Long.fromNumber(number);
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}

export interface MessageFns<T> {
  encode(message: T, writer?: BinaryWriter): BinaryWriter;
  decode(input: BinaryReader | Uint8Array, length?: number): T;
  fromJSON(object: any): T;
  toJSON(message: T): unknown;
  create(base?: DeepPartial<T>): T;
  fromPartial(object: DeepPartial<T>): T;
}
