'use client'

import { TTweet } from "@/types";
import { TweetCard } from "@/components/tweet/Card";

type Props = {
  tweets: TTweet[];
}

export const TweetWindow = ({ tweets }: Props) => {
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

