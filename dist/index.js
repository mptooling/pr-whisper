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
Object.defineProperty(exports, "__esModule", { value: true });
const core_1 = require("@octokit/core");
// @ts-ignore
const process = __importStar(require("process"));
// Instantiate Octokit with authentication
const octokit = new core_1.Octokit({
    auth: process.env.GITHUB_TOKEN, // Set this as an environment variable in your repo's settings
});
async function getChangedFiles(owner, repo, pullNumber) {
    // GitHub API: Get the list of changed files in the pull request
    const response = await octokit.request('GET /repos/{owner}/{repo}/pulls/{pull_number}/files', {
        owner,
        repo,
        pull_number: pullNumber,
    });
    return response.data.map((file) => file.filename);
}
async function createReview(owner, repo, pullNumber, commitId, comments, body, event) {
    // GitHub API: Create a review on the pull request
    await octokit.request('POST /repos/{owner}/{repo}/pulls/{pull_number}/reviews', {
        owner,
        repo,
        pull_number: pullNumber,
        commit_id: commitId,
        event,
        body,
        comments,
    });
}
async function checkTsFiles(files, owner, repo, pullNumber, commitId) {
    files.forEach(function (file) {
        console.log(file);
    });
    await createReview(owner, repo, pullNumber, commitId, [], "TEST.", "COMMENT");
}
async function run() {
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
    const changedFiles = await getChangedFiles(owner, repo, pullNumber);
    // Check TypeScript files and submit a review based on the results
    await checkTsFiles(changedFiles, owner, repo, pullNumber, commitId);
}
run().catch(error => {
    console.error(error);
    process.exit(1);
});
