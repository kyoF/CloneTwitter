export type TTweet = {
  tweetId: string;
  userId: string;
  text: string;
  likesCount: number;
}

export type TUser = {
  userId: string;
  name: string;
  tweets: TTweet[];
}

