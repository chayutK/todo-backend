import { test, expect } from "@playwright/test";

test("should response one todo when request POST /api/v1/todos", async ({
	request,
}) => {
	const reps = await request.post("http://localhost:8080/api/v1/todos", {
		data: {
			Title: "Go Tutorial",
			Status: "active",
		},
	});

	expect(reps.ok).toBeTruthy();
	const response = await reps.json();
	expect(response).toEqual(
		expect.objectContaining({
			ID: expect.any(Number),
			Title: "Go Tutorial",
			Status: "active",
		})
	);

	const id = response["ID"];
	await request.delete("http://localhost:8080/api/v1/todos/" + String(id));
});

test("should response with the same todo as input when request POST /api/v1/todos and GET /api/v1/todos", async ({
	request,
}) => {
	const title = "Go tutorial";
	const status = "active";

	const reps = await request.post("http://localhost:8080/api/v1/todos", {
		data: {
			Title: title,
			Status: status,
		},
	});

	expect(reps.ok).toBeTruthy();
	const response = await reps.json();
	const id = response["ID"];
	const getReps = await request.get(
		"http://localhost:8080/api/v1/todos/" + String(id)
	);

	expect(getReps.ok).toBeTruthy();

	const getResponse = await getReps.json();
	expect(getResponse).toEqual(
		expect.objectContaining({
			ID: expect.any(Number),
			Title: title,
			Status: status,
		})
	);

	await request.delete("http://localhost:8080/api/v1/todos/" + String(id));
});

test("should response with same all todo as input when request POST /api/v1/todos and GET /api/v1/todos", async ({
	request,
}) => {
	const title = "Go tutorial";
	const status = "active";
	const title2 = "Java tutorial";
	const status2 = "active";
	const title3 = "C tutorial";
	const status3 = "inactive";

	const reps1 = await request.post("http://localhost:8080/api/v1/todos", {
		data: {
			Title: title,
			Status: status,
		},
	});
	const reps2 = await request.post("http://localhost:8080/api/v1/todos", {
		data: {
			Title: title2,
			Status: status2,
		},
	});
	const reps3 = await request.post("http://localhost:8080/api/v1/todos", {
		data: {
			Title: title3,
			Status: status3,
		},
	});

	expect(reps1.ok).toBeTruthy();
	expect(reps2.ok).toBeTruthy();
	expect(reps3.ok).toBeTruthy();
	const response1 = await reps1.json();
	const id1 = response1["ID"];
	const response2 = await reps2.json();
	const id2 = response2["ID"];
	const response3 = await reps3.json();
	const id3 = response3["ID"];

	const getReps = await request.get("http://localhost:8080/api/v1/todos");

	expect(getReps.ok).toBeTruthy();

	expect(await getReps.json()).toEqual(
		expect.arrayContaining([
			{
				ID: id1,
				Title: title,
				Status: status,
			},
			{
				ID: id2,
				Title: title2,
				Status: status2,
			},
			{
				ID: id3,
				Title: title3,
				Status: status3,
			},
		])
	);

	await request.delete("http://localhost:8080/api/v1/todos/" + String(id1));
	await request.delete("http://localhost:8080/api/v1/todos/" + String(id2));
	await request.delete("http://localhost:8080/api/v1/todos/" + String(id3));
});

test('should response with "success" with DELETE /api/v1/todos/:id and must not be able to find todo with that id', async ({
	request,
}) => {
	const reps = await request.post("http://localhost:8080/api/v1/todos", {
		data: {
			Title: "Go Tutorial",
			Status: "active",
		},
	});

	expect(reps.ok).toBeTruthy();
	const response = await reps.json();

	const id = response["ID"];
	const deleteReps = await request.delete(
		"http://localhost:8080/api/v1/todos/" + String(id)
	);
	expect(deleteReps.ok).toBeTruthy();

	expect(await deleteReps.json()).toEqual("successful");

	const getReps = await request.get(
		"http://localhost:8080/api/v1/todos/" + String(id)
	);

	expect(getReps.ok).toBeTruthy();
	expect(await getReps.json()).toEqual(
		expect.objectContaining({
			ID: 0,
			Title: "",
			Status: "",
		})
	);
});

test("should response with new update row when request PUT /api/v1/todos/:id", async ({
	request,
}) => {
	const title = "Go tutorial";
	const status = "active";
	const reps = await request.post("http://localhost:8080/api/v1/todos", {
		data: {
			Title: title,
			Status: status,
		},
	});

	expect(reps.ok).toBeTruthy();
	const getRes = await reps.json();
	const id = getRes["ID"];

	const newtitle = "Go Tutorial V2";
	const newstatus = "inactive";
	const updateReps = await request.put(
		"http://localhost:8080/api/v1/todos/" + String(id),
		{
			data: {
				Title: newtitle,
				Status: newstatus,
			},
		}
	);

	const putRes = await updateReps.json();

	expect(updateReps.ok()).toBeTruthy();
	expect(putRes).toEqual(
		expect.objectContaining({
			ID: id,
			Title: newtitle,
			Status: newstatus,
		})
	);

	await request.delete("http://localhost:8080/api/v1/todos/" + String(id));
});

test("should response with new update row when request PATCH /api/v1/todos/:id/action/status", async ({
	request,
}) => {
	const title = "Go tutorial";
	const status = "active";
	const reps = await request.post("http://localhost:8080/api/v1/todos", {
		data: {
			Title: title,
			Status: status,
		},
	});
	expect(reps.ok).toBeTruthy();

	const getRes = await reps.json();
	const id = getRes["ID"];
	const newstatus = "inactive";
	const updateReps = await request.patch(
		"http://localhost:8080/api/v1/todos/" + String(id) + "/action/status",
		{
			data: {
				Status: newstatus,
			},
		}
	);
	expect(updateReps.ok()).toBeTruthy();

	await request.delete("http://localhost:8080/api/v1/todos/" + String(id));
});

test("should response with new update row when request PATCH /api/v1/todos/:id/action/title", async ({
	request,
}) => {
	const title = "Go tutorial";
	const status = "active";
	const reps = await request.post("http://localhost:8080/api/v1/todos", {
		data: {
			Title: title,
			Status: status,
		},
	});

	expect(reps.ok).toBeTruthy();

	const getRes = await reps.json();
	const id = getRes["ID"];
	const newtitle = "Go Tutorial V2";

	const updateReps = await request.patch(
		"http://localhost:8080/api/v1/todos/" + String(id) + "/action/title",
		{
			data: {
				Title: newtitle,
			},
		}
	);

	expect(updateReps.ok()).toBeTruthy();

	await request.delete("http://localhost:8080/api/v1/todos/" + String(id));
});
