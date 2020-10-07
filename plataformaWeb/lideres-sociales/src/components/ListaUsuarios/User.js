import React from "react";

export default function User(props) {
  const { user } = props;

  return (
    <div>
      <h2>{user.email}</h2>
      <h2>{user.role}</h2>
    </div>
  );
}
