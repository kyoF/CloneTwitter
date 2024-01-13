import type { TUser } from '@/types';
import { TweetWindow } from '@/components/TweetWindow';
import { TweetButton } from '@/components/TweetButton';

export default async function User({ params }: { params: { userId: string }}) {
  const res: Response = await fetch(
    `${process.env.API_URL}/user/${params.userId}`,
    { cache: 'no-store' }
  );
  const user: TUser = await res.json();

  return (
    <>
      <div>{user.userId}</div>
      <div>{user.name}</div>
      <TweetWindow tweets={user.tweets} />
      <TweetButton userId={params.userId} />
    </>
  );
}

