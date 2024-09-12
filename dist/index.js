"use strict";
var __createBinding = (this && this.__createBinding) || (Object.create ? (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    var desc = Object.getOwnPropertyDescriptor(m, k);
    if (!desc || ("get" in desc ? !m.__esModule : desc.writable || desc.configurable)) {
      desc = { enumerable: true, get: function() { return m[k]; } };
    }
    Object.defineProperty(o, k2, desc);
}) : (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    o[k2] = m[k];
}));
var __setModuleDefault = (this && this.__setModuleDefault) || (Object.create ? (function(o, v) {
    Object.defineProperty(o, "default", { enumerable: true, value: v });
}) : function(o, v) {
    o["default"] = v;
});
var __importStar = (this && this.__importStar) || function (mod) {
    if (mod && mod.__esModule) return mod;
    var result = {};
    if (mod != null) for (var k in mod) if (k !== "default" && Object.prototype.hasOwnProperty.call(mod, k)) __createBinding(result, mod, k);
    __setModuleDefault(result, mod);
    return result;
};
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
Object.defineProperty(exports, "__esModule", { value: true });
const rest_1 = require("@octokit/rest");
const process = __importStar(require("process"));
const octokit = new rest_1.Octokit({
    auth: process.env.GITHUB_TOKEN, // Set this as an environment variable in your repo's settings
});
function getChangedFiles(owner, repo, pullNumber) {
    return __awaiter(this, void 0, void 0, function* () {
        const { data } = yield octokit.pulls.listFiles({
            owner,
            repo,
            pull_number: pullNumber,
        });
        return data.map(file => file.filename);
    });
}
function createReview(owner, repo, pullNumber, commitId, comments, body, event) {
    return __awaiter(this, void 0, void 0, function* () {
        yield octokit.pulls.createReview({
            owner,
            repo,
            pull_number: pullNumber,
            commit_id: commitId,
            event, // "COMMENT", "APPROVE", or "REQUEST_CHANGES"
            body, // Overall review message (optional)
            comments, // The array of comments
        });
    });
}
function checkTsFiles(files, owner, repo, pullNumber, commitId) {
    return __awaiter(this, void 0, void 0, function* () {
        files.forEach(function (file) {
            console.log(file);
        });
        yield createReview(owner, repo, pullNumber, commitId, [], "TEST.", "COMMENT");
    });
}
function run() {
    return __awaiter(this, void 0, void 0, function* () {
        var _a, _b;
        const owner = (_a = process.env.GITHUB_REPOSITORY) === null || _a === void 0 ? void 0 : _a.split("/")[0];
        const repo = (_b = process.env.GITHUB_REPOSITORY) === null || _b === void 0 ? void 0 : _b.split("/")[1];
        const pullNumber = parseInt(process.env.GITHUB_PULL_REQUEST_NUMBER || "0", 10);
        const commitId = process.env.GITHUB_SHA; // Get the commit SHA from the environment variable
        if (!owner || !repo || !pullNumber || !commitId) {
            console.error("Required environment variables are missing.");
            process.exit(1);
        }
        // Get the list of changed files in the PR
        const changedFiles = yield getChangedFiles(owner, repo, pullNumber);
        // Check TypeScript files and submit a review based on the results
        yield checkTsFiles(changedFiles, owner, repo, pullNumber, commitId);
    });
}
run().catch(error => {
    console.error(error);
    process.exit(1);
});
