"use strict";

const { test, expect, afterEach } = require("@jest/globals");

// The logger captures the real console methods when it is required, so spy on
// console first and load a fresh copy of the module per test.
function loadLogger(level) {
  const spies = {
    debug: jest.spyOn(console, "debug").mockImplementation(() => {}),
    info: jest.spyOn(console, "info").mockImplementation(() => {}),
    warn: jest.spyOn(console, "warn").mockImplementation(() => {}),
    error: jest.spyOn(console, "error").mockImplementation(() => {}),
  };
  let createLogger;
  jest.isolateModules(() => {
    ({ createLogger } = require("../src/logger"));
  });
  return { logger: createLogger(level), spies };
}

afterEach(() => {
  jest.restoreAllMocks();
});

// trace and debug share the console.debug sink, and error and fatal share
// console.error, so identify what was emitted by message.
const MESSAGES = {
  trace: "trace message",
  debug: "debug message",
  info: "info message",
  warn: "warn message",
  error: "error message",
  fatal: "fatal message",
};

function emitEveryLevel(logger) {
  Object.entries(MESSAGES).forEach(([level, message]) =>
    logger[level](message),
  );
}

function emittedMessages(spies) {
  const messages = [];
  ["debug", "info", "warn", "error"].forEach((sink) => {
    spies[sink].mock.calls.forEach(([message]) => messages.push(message));
  });
  return messages;
}

function expectEmitted(level, expectedLevels) {
  const { logger, spies } = loadLogger(level);
  emitEveryLevel(logger);
  expect(emittedMessages(spies).sort()).toEqual(
    expectedLevels.map((name) => MESSAGES[name]).sort(),
  );
}

test("all enables every level", () => {
  expectEmitted("all", ["trace", "debug", "info", "warn", "error", "fatal"]);
});

test("trace enables every level", () => {
  expectEmitted("trace", ["trace", "debug", "info", "warn", "error", "fatal"]);
});

test("debug suppresses trace", () => {
  expectEmitted("debug", ["debug", "info", "warn", "error", "fatal"]);
});

test("info suppresses trace and debug", () => {
  expectEmitted("info", ["info", "warn", "error", "fatal"]);
});

test("warn suppresses trace, debug, and info", () => {
  expectEmitted("warn", ["warn", "error", "fatal"]);
});

test("error suppresses everything below error", () => {
  expectEmitted("error", ["error", "fatal"]);
});

test("fatal suppresses everything below fatal", () => {
  expectEmitted("fatal", ["fatal"]);
});

test("an unknown level falls back to the default (warn)", () => {
  expectEmitted("nope", ["warn", "error", "fatal"]);
});

test("the default level is warn when no level is given", () => {
  expectEmitted(undefined, ["warn", "error", "fatal"]);
});

test("log is an alias of info", () => {
  const { logger, spies } = loadLogger("info");
  logger.log("log message");
  logger.debug("debug message");
  expect(spies.info).toHaveBeenCalledWith("log message");
  expect(spies.debug).not.toHaveBeenCalled();
});

test("setLevel still reaches the native console after global.console is replaced", () => {
  const { logger, spies } = loadLogger("warn");
  const originalConsole = global.console;
  global.console = logger;
  try {
    logger.setLevel("info");
    logger.trace("trace message");
    logger.info("info message");
    logger.warn("warn message");
  } finally {
    global.console = originalConsole;
  }

  expect(spies.debug).not.toHaveBeenCalled();
  expect(spies.info).toHaveBeenCalledWith("info message");
  expect(spies.warn).toHaveBeenCalledWith("warn message");
});
