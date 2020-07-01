package counter

import "module06/app/services/post"

func PostCount(client post.Post) (int, error) {
	posts, err := client.GetList()
	if err != nil {
		return 0, nil
	}

	return len(posts), nil
}
