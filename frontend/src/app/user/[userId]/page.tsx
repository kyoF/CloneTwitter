import { UserContainer } from "@/components/user/Container";

export default async function UserPage({ params }: { params: { userId: string } }) {
  return (
    <>
      <UserContainer userId={params.userId} />
    </>
  );
}

