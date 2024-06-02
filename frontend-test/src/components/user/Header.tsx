import { TUser } from "@/types"

export const UserHeader = ({ user }: { user: TUser }) => {
  return (
    <>
      <div>============</div>
      <div>{user.userId}</div>
      <div>{user.name}</div>
      <div>{user.profile}</div>
      <div>============</div>
    </>
  )
}
