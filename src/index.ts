import { Octokit } from "@octokit/rest";
// @ts-ignore
import * as process from "process";

const octokit = new Octokit({
    auth: process.env.GITHUB_TOKEN, // Set this as an environment variable in your repo's settings
});

async function getChangedFiles(owner: string, repo: string, pullNumber: number) {
    const { data } = await octokit.pulls.listFiles({
        owner,
        repo,
        pull_number: pullNumber,
    });

    return data.map(file => file.filename);
}

async function createReview(
    owner: string,
    repo: string,
    pullNumber: number,
    commitId: string,
    comments: Array<{ path: string, position: number, body: string }>,
    body: string,
    event: "COMMENT" | "APPROVE" | "REQUEST_CHANGES"
) {
    await octokit.pulls.createReview({
        owner,
        repo,
        pull_number: pullNumber,
        commit_id: commitId,
        event, // "COMMENT", "APPROVE", or "REQUEST_CHANGES"
        body, // Overall review message (optional)
        comments, // The array of comments
    });
}

async function checkTsFiles(
    files: string[],
    owner: string,
    repo: string,
    pullNumber: number,
    commitId: string
) {
    files.forEach(function (file){
        console.log(file)
    })
    await createReview(owner, repo, pullNumber, commitId, [], "TEST.", "COMMENT");
}

async function run() {
    const owner = process.env.GITHUB_REPOSITORY?.split("/")[0];
    const repo = process.env.GITHUB_REPOSITORY?.split("/")[1];
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
