import type { TTweet, TUser } from "@/types";
import { UserHeader } from "@/components/user/Header";
import { TweetContainer } from "@/components/tweet/Container"

export const UserContainer = async ({ userId }: { userId: string }) => {
  const res: Response = await fetch(
    `${process.env.API_URL}/user/${userId}/tweets`,
    { cache: 'no-store' }
  );
  const tweets: TTweet[] = await res.json();
  console.log(tweets)
  console.log(`${process.env.API_URL}/user/${userId}/tweets`);
  console.log(res);

  const res2: Response = await fetch(
    `${process.env.API_URL}/user/${userId}`,
    { cache: 'no-store' }
  );
  const user: TUser = await res2.json();
  console.log(user)

  return (
  <>
    <UserHeader user={user} />
    <TweetContainer tweets={tweets}/>
    </>
  );
}
