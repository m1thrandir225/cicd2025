import { apiUrl } from "~/api/config";
import * as z from "zod";

const refreshTokenSchema = z.object({
  refresh_token: z.string(),
  user_id: z.string(),
});

export type RefreshTokenResponse = {
  access_token: string;
  access_token_expires_at: string;
};

export default defineEventHandler(async (event) => {
  const body = await readValidatedBody(event, refreshTokenSchema.safeParse);

  if (!body.success) {
    throw body.error.issues;
  }

  const response = await fetch(`${apiUrl}/refresh-token`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(body.data),
  });

  const data = await response.json();

  if (!response.ok) {
    throw createError({
      status: response.status,
      statusMessage: response.statusText,
      message: data.error,
    });
  }

  return data as RefreshTokenResponse;
});
