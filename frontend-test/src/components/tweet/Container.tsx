'use client'

import { TTweet } from "@/types";
import { TweetCard } from "@/components/tweet/Card";

export const TweetContainer = ({ tweets }: { tweets: TTweet[] }) => {
  return (
    <>
      {tweets.map((tweet: TTweet) => (
        <div key={tweet.tweetId}>
          <TweetCard {...tweet} />
        </div>
      ))}
    </>
  );
}

