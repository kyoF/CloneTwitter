import type { TTweet, TUser } from '@/types';
import { UserHeader } from '@/components/user/Header';
import { UserBody } from '@/components/user/Body';

export default async function UserPage({ params }: { params: { userId: string } }) {
  const res: Response = await fetch(
    `${process.env.API_URL}/user/${params.userId}/tweets`,
    { cache: 'no-store' }
  );
  const tweets: TTweet[] = await res.json();

  const res2: Response = await fetch(
    `${process.env.API_URL}/user/${params.userId}`,
    { cache: 'no-store' }
  );
  const user: TUser = await res2.json();
  return (
    <>
      <UserHeader user={user} />
      <UserBody tweets={tweets} />
    </>
  );
}

