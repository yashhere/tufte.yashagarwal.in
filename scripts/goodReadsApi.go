package main

import "encoding/xml"

type GoodreadsUser struct {
	XMLName xml.Name `xml:"GoodreadsResponse"`
	Text    string   `xml:",chardata"`
	Request struct {
		Text           string `xml:",chardata"`
		Authentication string `xml:"authentication"`
		Key            string `xml:"key"`
		Method         string `xml:"method"`
	} `xml:"Request"`
	User struct {
		Text            string `xml:",chardata"`
		ID              string `xml:"id"`
		Name            string `xml:"name"`
		UserName        string `xml:"user_name"`
		Link            string `xml:"link"`
		ImageURL        string `xml:"image_url"`
		SmallImageURL   string `xml:"small_image_url"`
		About           string `xml:"about"`
		Age             string `xml:"age"`
		Gender          string `xml:"gender"`
		Location        string `xml:"location"`
		Website         string `xml:"website"`
		Joined          string `xml:"joined"`
		LastActive      string `xml:"last_active"`
		Interests       string `xml:"interests"`
		FavoriteBooks   string `xml:"favorite_books"`
		FavoriteAuthors string `xml:"favorite_authors"`
		UpdatesRssURL   string `xml:"updates_rss_url"`
		ReviewsRssURL   string `xml:"reviews_rss_url"`
		FriendsCount    struct {
			Text string `xml:",chardata"`
			Type string `xml:"type,attr"`
		} `xml:"friends_count"`
		GroupsCount  string `xml:"groups_count"`
		ReviewsCount struct {
			Text string `xml:",chardata"`
			Type string `xml:"type,attr"`
		} `xml:"reviews_count"`
		UserShelves struct {
			Text      string `xml:",chardata"`
			Type      string `xml:"type,attr"`
			UserShelf []struct {
				Text string `xml:",chardata"`
				ID   struct {
					Text string `xml:",chardata"`
					Type string `xml:"type,attr"`
				} `xml:"id"`
				Name      string `xml:"name"`
				BookCount struct {
					Text string `xml:",chardata"`
					Type string `xml:"type,attr"`
				} `xml:"book_count"`
				ExclusiveFlag struct {
					Text string `xml:",chardata"`
					Type string `xml:"type,attr"`
				} `xml:"exclusive_flag"`
				Sort struct {
					Text string `xml:",chardata"`
					Nil  string `xml:"nil,attr"`
				} `xml:"sort"`
				Order struct {
					Text string `xml:",chardata"`
					Nil  string `xml:"nil,attr"`
				} `xml:"order"`
				PerPage struct {
					Text string `xml:",chardata"`
					Type string `xml:"type,attr"`
					Nil  string `xml:"nil,attr"`
				} `xml:"per_page"`
				DisplayFields string `xml:"display_fields"`
				Featured      struct {
					Text string `xml:",chardata"`
					Type string `xml:"type,attr"`
				} `xml:"featured"`
				RecommendFor struct {
					Text string `xml:",chardata"`
					Type string `xml:"type,attr"`
				} `xml:"recommend_for"`
				Sticky struct {
					Text string `xml:",chardata"`
					Type string `xml:"type,attr"`
					Nil  string `xml:"nil,attr"`
				} `xml:"sticky"`
			} `xml:"user_shelf"`
		} `xml:"user_shelves"`
		Updates struct {
			Text   string `xml:",chardata"`
			Type   string `xml:"type,attr"`
			Update []struct {
				Text       string `xml:",chardata"`
				Type       string `xml:"type,attr"`
				ActionText string `xml:"action_text"`
				Link       string `xml:"link"`
				ImageURL   string `xml:"image_url"`
				Actor      struct {
					Text string `xml:",chardata"`
					ID   struct {
						Text string `xml:",chardata"`
						Type string `xml:"type,attr"`
					} `xml:"id"`
					Name     string `xml:"name"`
					ImageURL string `xml:"image_url"`
					Link     string `xml:"link"`
				} `xml:"actor"`
				UpdatedAt string `xml:"updated_at"`
				Action    struct {
					Text   string `xml:",chardata"`
					Type   string `xml:"type,attr"`
					Rating string `xml:"rating"`
				} `xml:"action"`
				Object struct {
					Text string `xml:",chardata"`
					Book struct {
						Text    string `xml:",chardata"`
						ID      string `xml:"id"`
						Title   string `xml:"title"`
						Link    string `xml:"link"`
						Authors struct {
							Text   string `xml:",chardata"`
							Author struct {
								Text     string `xml:",chardata"`
								ID       string `xml:"id"`
								Name     string `xml:"name"`
								Role     string `xml:"role"`
								ImageURL struct {
									Text    string `xml:",chardata"`
									Nophoto string `xml:"nophoto,attr"`
								} `xml:"image_url"`
								SmallImageURL struct {
									Text    string `xml:",chardata"`
									Nophoto string `xml:"nophoto,attr"`
								} `xml:"small_image_url"`
								Link             string `xml:"link"`
								AverageRating    string `xml:"average_rating"`
								RatingsCount     string `xml:"ratings_count"`
								TextReviewsCount string `xml:"text_reviews_count"`
							} `xml:"author"`
						} `xml:"authors"`
					} `xml:"book"`
					ReadStatus struct {
						Text string `xml:",chardata"`
						ID   struct {
							Text string `xml:",chardata"`
							Type string `xml:"type,attr"`
						} `xml:"id"`
						ReviewID struct {
							Text string `xml:",chardata"`
							Type string `xml:"type,attr"`
						} `xml:"review_id"`
						UserID struct {
							Text string `xml:",chardata"`
							Type string `xml:"type,attr"`
						} `xml:"user_id"`
						OldStatus struct {
							Text string `xml:",chardata"`
							Nil  string `xml:"nil,attr"`
						} `xml:"old_status"`
						Status    string `xml:"status"`
						UpdatedAt struct {
							Text string `xml:",chardata"`
							Type string `xml:"type,attr"`
						} `xml:"updated_at"`
						Review struct {
							Text string `xml:",chardata"`
							ID   struct {
								Text string `xml:",chardata"`
								Type string `xml:"type,attr"`
							} `xml:"id"`
							UserID struct {
								Text string `xml:",chardata"`
								Type string `xml:"type,attr"`
							} `xml:"user_id"`
							BookID struct {
								Text string `xml:",chardata"`
								Type string `xml:"type,attr"`
							} `xml:"book_id"`
							Rating struct {
								Text string `xml:",chardata"`
								Type string `xml:"type,attr"`
							} `xml:"rating"`
							ReadStatus string `xml:"read_status"`
							SellFlag   struct {
								Text string `xml:",chardata"`
								Type string `xml:"type,attr"`
							} `xml:"sell_flag"`
							Review struct {
								Text string `xml:",chardata"`
								Nil  string `xml:"nil,attr"`
							} `xml:"review"`
							Recommendation struct {
								Text string `xml:",chardata"`
								Nil  string `xml:"nil,attr"`
							} `xml:"recommendation"`
							ReadAt struct {
								Text string `xml:",chardata"`
								Type string `xml:"type,attr"`
								Nil  string `xml:"nil,attr"`
							} `xml:"read_at"`
							UpdatedAt struct {
								Text string `xml:",chardata"`
								Type string `xml:"type,attr"`
							} `xml:"updated_at"`
							CreatedAt struct {
								Text string `xml:",chardata"`
								Type string `xml:"type,attr"`
							} `xml:"created_at"`
							CommentsCount struct {
								Text string `xml:",chardata"`
								Type string `xml:"type,attr"`
							} `xml:"comments_count"`
							Weight struct {
								Text string `xml:",chardata"`
								Type string `xml:"type,attr"`
							} `xml:"weight"`
							RatingsSum struct {
								Text string `xml:",chardata"`
								Type string `xml:"type,attr"`
							} `xml:"ratings_sum"`
							RatingsCount struct {
								Text string `xml:",chardata"`
								Type string `xml:"type,attr"`
							} `xml:"ratings_count"`
							SpoilerFlag struct {
								Text string `xml:",chardata"`
								Type string `xml:"type,attr"`
							} `xml:"spoiler_flag"`
							RecommenderUserId1 struct {
								Text string `xml:",chardata"`
								Type string `xml:"type,attr"`
							} `xml:"recommender_user_id1"`
							RecommenderUserName1 struct {
								Text string `xml:",chardata"`
								Nil  string `xml:"nil,attr"`
							} `xml:"recommender_user_name1"`
							WorkID struct {
								Text string `xml:",chardata"`
								Type string `xml:"type,attr"`
							} `xml:"work_id"`
							LastCommentAt struct {
								Text string `xml:",chardata"`
								Type string `xml:"type,attr"`
								Nil  string `xml:"nil,attr"`
							} `xml:"last_comment_at"`
							StartedAt struct {
								Text string `xml:",chardata"`
								Type string `xml:"type,attr"`
								Nil  string `xml:"nil,attr"`
							} `xml:"started_at"`
							HiddenFlag struct {
								Text string `xml:",chardata"`
								Type string `xml:"type,attr"`
							} `xml:"hidden_flag"`
							LanguageCode struct {
								Text string `xml:",chardata"`
								Type string `xml:"type,attr"`
								Nil  string `xml:"nil,attr"`
							} `xml:"language_code"`
							LastRevisionAt struct {
								Text string `xml:",chardata"`
								Type string `xml:"type,attr"`
							} `xml:"last_revision_at"`
							NonFriendsRatingCount struct {
								Text string `xml:",chardata"`
								Type string `xml:"type,attr"`
							} `xml:"non_friends_rating_count"`
							EncryptedNotes struct {
								Text string `xml:",chardata"`
								Nil  string `xml:"nil,attr"`
							} `xml:"encrypted_notes"`
							BookURI string `xml:"book_uri"`
							Notes   struct {
								Text string `xml:",chardata"`
								Nil  string `xml:"nil,attr"`
							} `xml:"notes"`
							Book struct {
								Text string `xml:",chardata"`
								ID   struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr"`
								} `xml:"id"`
								WorkID struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr"`
								} `xml:"work_id"`
								ISBN struct {
									Text string `xml:",chardata"`
									Nil  string `xml:"nil,attr"`
								} `xml:"isbn"`
								Isbn13 struct {
									Text string `xml:",chardata"`
									Nil  string `xml:"nil,attr"`
								} `xml:"isbn13"`
								Title       string `xml:"title"`
								SortByTitle string `xml:"sort_by_title"`
								AuthorID    struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr"`
								} `xml:"author_id"`
								AuthorRole struct {
									Text string `xml:",chardata"`
									Nil  string `xml:"nil,attr"`
								} `xml:"author_role"`
								Asin struct {
									Text string `xml:",chardata"`
									Nil  string `xml:"nil,attr"`
								} `xml:"asin"`
								Description struct {
									Text string `xml:",chardata"`
									Nil  string `xml:"nil,attr"`
								} `xml:"description"`
								Format          string `xml:"format"`
								PublicationYear struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr"`
								} `xml:"publication_year"`
								PublicationMonth struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr"`
									Nil  string `xml:"nil,attr"`
								} `xml:"publication_month"`
								PublicationDay struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr"`
									Nil  string `xml:"nil,attr"`
								} `xml:"publication_day"`
								NumPages struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr"`
								} `xml:"num_pages"`
								Publisher          string `xml:"publisher"`
								LanguageCode       string `xml:"language_code"`
								EditionInformation struct {
									Text string `xml:",chardata"`
									Nil  string `xml:"nil,attr"`
								} `xml:"edition_information"`
								URL struct {
									Text string `xml:",chardata"`
									Nil  string `xml:"nil,attr"`
								} `xml:"url"`
								SourceURL struct {
									Text string `xml:",chardata"`
									Nil  string `xml:"nil,attr"`
								} `xml:"source_url"`
								ImageUploadedAt struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr"`
								} `xml:"image_uploaded_at"`
								S3ImageAt struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr"`
								} `xml:"s3_image_at"`
								ReviewsCount struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr"`
								} `xml:"reviews_count"`
								RatingsSum struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr"`
								} `xml:"ratings_sum"`
								RatingsCount struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr"`
								} `xml:"ratings_count"`
								TextReviewsCount struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr"`
								} `xml:"text_reviews_count"`
								BookAuthorsCount struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr"`
								} `xml:"book_authors_count"`
								UpdatedAt struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr"`
								} `xml:"updated_at"`
								CreatedAt struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr"`
								} `xml:"created_at"`
								AuthorIDUpdaterUserID struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr"`
								} `xml:"author_id_updater_user_id"`
								AuthorRoleUpdaterUserID struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr"`
									Nil  string `xml:"nil,attr"`
								} `xml:"author_role_updater_user_id"`
								DescriptionUpdaterUserID struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr"`
								} `xml:"description_updater_user_id"`
								EditionInformationUpdaterUserID struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr"`
									Nil  string `xml:"nil,attr"`
								} `xml:"edition_information_updater_user_id"`
								FormatUpdaterUserID struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr"`
								} `xml:"format_updater_user_id"`
								ImageUpdaterUserID struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr"`
								} `xml:"image_updater_user_id"`
								ISBNUpdaterUserID struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr"`
								} `xml:"isbn_updater_user_id"`
								Isbn13UpdaterUserID struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr"`
									Nil  string `xml:"nil,attr"`
								} `xml:"isbn13_updater_user_id"`
								LanguageUpdaterUserID struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr"`
								} `xml:"language_updater_user_id"`
								NumPagesUpdaterUserID struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr"`
								} `xml:"num_pages_updater_user_id"`
								PublicationDateUpdaterUserID struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr"`
								} `xml:"publication_date_updater_user_id"`
								PublisherUpdaterUserID struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr"`
								} `xml:"publisher_updater_user_id"`
								TitleUpdaterUserID struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr"`
								} `xml:"title_updater_user_id"`
								URLUpdaterUserID struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr"`
									Nil  string `xml:"nil,attr"`
								} `xml:"url_updater_user_id"`
								TitleLanguageCode struct {
									Text string `xml:",chardata"`
									Nil  string `xml:"nil,attr"`
								} `xml:"title_language_code"`
								PublisherLanguageCode struct {
									Text string `xml:",chardata"`
									Nil  string `xml:"nil,attr"`
								} `xml:"publisher_language_code"`
								DescriptionLanguageCode struct {
									Text string `xml:",chardata"`
									Nil  string `xml:"nil,attr"`
								} `xml:"description_language_code"`
								AsinUpdaterUserID struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr"`
									Nil  string `xml:"nil,attr"`
								} `xml:"asin_updater_user_id"`
								BookURI string `xml:"book_uri"`
								Author  struct {
									Text string `xml:",chardata"`
									ID   struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr"`
									} `xml:"id"`
									Name      string `xml:"name"`
									UpdatedAt struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr"`
									} `xml:"updated_at"`
									CreatedAt struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr"`
									} `xml:"created_at"`
									ImageUploadedAt struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr"`
									} `xml:"image_uploaded_at"`
									UserID struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr"`
										Nil  string `xml:"nil,attr"`
									} `xml:"user_id"`
									Gender      string `xml:"gender"`
									CountryCode string `xml:"country_code"`
									BornAt      struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr"`
									} `xml:"born_at"`
									DiedAt struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr"`
										Nil  string `xml:"nil,attr"`
									} `xml:"died_at"`
									About          string `xml:"about"`
									UploaderUserID struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr"`
									} `xml:"uploader_user_id"`
									ImageCopyright string `xml:"image_copyright"`
									Influences     string `xml:"influences"`
									URL            string `xml:"url"`
									Genre1         string `xml:"genre1"`
									Genre2         string `xml:"genre2"`
									Genre3         string `xml:"genre3"`
									BooksCount     struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr"`
									} `xml:"books_count"`
									ReviewsCount struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr"`
									} `xml:"reviews_count"`
									RatingsSum struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr"`
									} `xml:"ratings_sum"`
									WorksCount struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr"`
									} `xml:"works_count"`
									Hometown   string `xml:"hometown"`
									RatingDist string `xml:"rating_dist"`
									S3ImageAt  struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr"`
										Nil  string `xml:"nil,attr"`
									} `xml:"s3_image_at"`
									RatingsCount struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr"`
									} `xml:"ratings_count"`
									TextReviewsCount struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr"`
									} `xml:"text_reviews_count"`
									AuthorProgramAt struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr"`
										Nil  string `xml:"nil,attr"`
									} `xml:"author_program_at"`
									BestBookID struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr"`
									} `xml:"best_book_id"`
									SortByName       string `xml:"sort_by_name"`
									ShelfDisplayName string `xml:"shelf_display_name"`
									AuthorURI        string `xml:"author_uri"`
								} `xml:"author"`
							} `xml:"book"`
						} `xml:"review"`
					} `xml:"read_status"`
					UserStatus struct {
						Text string `xml:",chardata"`
						ID   struct {
							Text string `xml:",chardata"`
							Type string `xml:"type,attr"`
						} `xml:"id"`
						UserID struct {
							Text string `xml:",chardata"`
							Type string `xml:"type,attr"`
						} `xml:"user_id"`
						BookID struct {
							Text string `xml:",chardata"`
							Type string `xml:"type,attr"`
						} `xml:"book_id"`
						Page struct {
							Text string `xml:",chardata"`
							Type string `xml:"type,attr"`
						} `xml:"page"`
						CommentsCount struct {
							Text string `xml:",chardata"`
							Type string `xml:"type,attr"`
						} `xml:"comments_count"`
						Body struct {
							Text string `xml:",chardata"`
							Nil  string `xml:"nil,attr"`
						} `xml:"body"`
						CreatedAt struct {
							Text string `xml:",chardata"`
							Type string `xml:"type,attr"`
						} `xml:"created_at"`
						UpdatedAt struct {
							Text string `xml:",chardata"`
							Type string `xml:"type,attr"`
						} `xml:"updated_at"`
						LastCommentAt struct {
							Text string `xml:",chardata"`
							Type string `xml:"type,attr"`
							Nil  string `xml:"nil,attr"`
						} `xml:"last_comment_at"`
						RatingsCount struct {
							Text string `xml:",chardata"`
							Type string `xml:"type,attr"`
						} `xml:"ratings_count"`
						Percent struct {
							Text string `xml:",chardata"`
							Type string `xml:"type,attr"`
						} `xml:"percent"`
						WorkID struct {
							Text string `xml:",chardata"`
							Type string `xml:"type,attr"`
						} `xml:"work_id"`
						NoteURI struct {
							Text string `xml:",chardata"`
							Nil  string `xml:"nil,attr"`
						} `xml:"note_uri"`
						NoteUpdatedAt struct {
							Text string `xml:",chardata"`
							Type string `xml:"type,attr"`
							Nil  string `xml:"nil,attr"`
						} `xml:"note_updated_at"`
						ReviewID struct {
							Text string `xml:",chardata"`
							Type string `xml:"type,attr"`
						} `xml:"review_id"`
						Book struct {
							Text string `xml:",chardata"`
							ID   struct {
								Text string `xml:",chardata"`
								Type string `xml:"type,attr"`
							} `xml:"id"`
							WorkID struct {
								Text string `xml:",chardata"`
								Type string `xml:"type,attr"`
							} `xml:"work_id"`
							ISBN struct {
								Text string `xml:",chardata"`
								Nil  string `xml:"nil,attr"`
							} `xml:"isbn"`
							Isbn13 struct {
								Text string `xml:",chardata"`
								Nil  string `xml:"nil,attr"`
							} `xml:"isbn13"`
							Title       string `xml:"title"`
							SortByTitle string `xml:"sort_by_title"`
							AuthorID    struct {
								Text string `xml:",chardata"`
								Type string `xml:"type,attr"`
							} `xml:"author_id"`
							AuthorRole struct {
								Text string `xml:",chardata"`
								Nil  string `xml:"nil,attr"`
							} `xml:"author_role"`
							Asin struct {
								Text string `xml:",chardata"`
								Nil  string `xml:"nil,attr"`
							} `xml:"asin"`
							Description struct {
								Text string `xml:",chardata"`
								Nil  string `xml:"nil,attr"`
							} `xml:"description"`
							Format          string `xml:"format"`
							PublicationYear struct {
								Text string `xml:",chardata"`
								Type string `xml:"type,attr"`
							} `xml:"publication_year"`
							PublicationMonth struct {
								Text string `xml:",chardata"`
								Type string `xml:"type,attr"`
								Nil  string `xml:"nil,attr"`
							} `xml:"publication_month"`
							PublicationDay struct {
								Text string `xml:",chardata"`
								Type string `xml:"type,attr"`
								Nil  string `xml:"nil,attr"`
							} `xml:"publication_day"`
							NumPages struct {
								Text string `xml:",chardata"`
								Type string `xml:"type,attr"`
							} `xml:"num_pages"`
							Publisher          string `xml:"publisher"`
							LanguageCode       string `xml:"language_code"`
							EditionInformation struct {
								Text string `xml:",chardata"`
								Nil  string `xml:"nil,attr"`
							} `xml:"edition_information"`
							URL struct {
								Text string `xml:",chardata"`
								Nil  string `xml:"nil,attr"`
							} `xml:"url"`
							SourceURL struct {
								Text string `xml:",chardata"`
								Nil  string `xml:"nil,attr"`
							} `xml:"source_url"`
							ImageUploadedAt struct {
								Text string `xml:",chardata"`
								Type string `xml:"type,attr"`
							} `xml:"image_uploaded_at"`
							S3ImageAt struct {
								Text string `xml:",chardata"`
								Type string `xml:"type,attr"`
							} `xml:"s3_image_at"`
							ReviewsCount struct {
								Text string `xml:",chardata"`
								Type string `xml:"type,attr"`
							} `xml:"reviews_count"`
							RatingsSum struct {
								Text string `xml:",chardata"`
								Type string `xml:"type,attr"`
							} `xml:"ratings_sum"`
							RatingsCount struct {
								Text string `xml:",chardata"`
								Type string `xml:"type,attr"`
							} `xml:"ratings_count"`
							TextReviewsCount struct {
								Text string `xml:",chardata"`
								Type string `xml:"type,attr"`
							} `xml:"text_reviews_count"`
							BookAuthorsCount struct {
								Text string `xml:",chardata"`
								Type string `xml:"type,attr"`
							} `xml:"book_authors_count"`
							UpdatedAt struct {
								Text string `xml:",chardata"`
								Type string `xml:"type,attr"`
							} `xml:"updated_at"`
							CreatedAt struct {
								Text string `xml:",chardata"`
								Type string `xml:"type,attr"`
							} `xml:"created_at"`
							AuthorIDUpdaterUserID struct {
								Text string `xml:",chardata"`
								Type string `xml:"type,attr"`
							} `xml:"author_id_updater_user_id"`
							AuthorRoleUpdaterUserID struct {
								Text string `xml:",chardata"`
								Type string `xml:"type,attr"`
								Nil  string `xml:"nil,attr"`
							} `xml:"author_role_updater_user_id"`
							DescriptionUpdaterUserID struct {
								Text string `xml:",chardata"`
								Type string `xml:"type,attr"`
							} `xml:"description_updater_user_id"`
							EditionInformationUpdaterUserID struct {
								Text string `xml:",chardata"`
								Type string `xml:"type,attr"`
								Nil  string `xml:"nil,attr"`
							} `xml:"edition_information_updater_user_id"`
							FormatUpdaterUserID struct {
								Text string `xml:",chardata"`
								Type string `xml:"type,attr"`
							} `xml:"format_updater_user_id"`
							ImageUpdaterUserID struct {
								Text string `xml:",chardata"`
								Type string `xml:"type,attr"`
							} `xml:"image_updater_user_id"`
							ISBNUpdaterUserID struct {
								Text string `xml:",chardata"`
								Type string `xml:"type,attr"`
							} `xml:"isbn_updater_user_id"`
							Isbn13UpdaterUserID struct {
								Text string `xml:",chardata"`
								Type string `xml:"type,attr"`
								Nil  string `xml:"nil,attr"`
							} `xml:"isbn13_updater_user_id"`
							LanguageUpdaterUserID struct {
								Text string `xml:",chardata"`
								Type string `xml:"type,attr"`
							} `xml:"language_updater_user_id"`
							NumPagesUpdaterUserID struct {
								Text string `xml:",chardata"`
								Type string `xml:"type,attr"`
							} `xml:"num_pages_updater_user_id"`
							PublicationDateUpdaterUserID struct {
								Text string `xml:",chardata"`
								Type string `xml:"type,attr"`
							} `xml:"publication_date_updater_user_id"`
							PublisherUpdaterUserID struct {
								Text string `xml:",chardata"`
								Type string `xml:"type,attr"`
							} `xml:"publisher_updater_user_id"`
							TitleUpdaterUserID struct {
								Text string `xml:",chardata"`
								Type string `xml:"type,attr"`
							} `xml:"title_updater_user_id"`
							URLUpdaterUserID struct {
								Text string `xml:",chardata"`
								Type string `xml:"type,attr"`
								Nil  string `xml:"nil,attr"`
							} `xml:"url_updater_user_id"`
							TitleLanguageCode struct {
								Text string `xml:",chardata"`
								Nil  string `xml:"nil,attr"`
							} `xml:"title_language_code"`
							PublisherLanguageCode struct {
								Text string `xml:",chardata"`
								Nil  string `xml:"nil,attr"`
							} `xml:"publisher_language_code"`
							DescriptionLanguageCode struct {
								Text string `xml:",chardata"`
								Nil  string `xml:"nil,attr"`
							} `xml:"description_language_code"`
							AsinUpdaterUserID struct {
								Text string `xml:",chardata"`
								Type string `xml:"type,attr"`
								Nil  string `xml:"nil,attr"`
							} `xml:"asin_updater_user_id"`
							BookURI string `xml:"book_uri"`
							Author  struct {
								Text string `xml:",chardata"`
								ID   struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr"`
								} `xml:"id"`
								Name      string `xml:"name"`
								UpdatedAt struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr"`
								} `xml:"updated_at"`
								CreatedAt struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr"`
								} `xml:"created_at"`
								ImageUploadedAt struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr"`
								} `xml:"image_uploaded_at"`
								UserID struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr"`
									Nil  string `xml:"nil,attr"`
								} `xml:"user_id"`
								Gender      string `xml:"gender"`
								CountryCode string `xml:"country_code"`
								BornAt      struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr"`
								} `xml:"born_at"`
								DiedAt struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr"`
									Nil  string `xml:"nil,attr"`
								} `xml:"died_at"`
								About          string `xml:"about"`
								UploaderUserID struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr"`
								} `xml:"uploader_user_id"`
								ImageCopyright string `xml:"image_copyright"`
								Influences     string `xml:"influences"`
								URL            string `xml:"url"`
								Genre1         string `xml:"genre1"`
								Genre2         string `xml:"genre2"`
								Genre3         string `xml:"genre3"`
								BooksCount     struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr"`
								} `xml:"books_count"`
								ReviewsCount struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr"`
								} `xml:"reviews_count"`
								RatingsSum struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr"`
								} `xml:"ratings_sum"`
								WorksCount struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr"`
								} `xml:"works_count"`
								Hometown   string `xml:"hometown"`
								RatingDist string `xml:"rating_dist"`
								S3ImageAt  struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr"`
									Nil  string `xml:"nil,attr"`
								} `xml:"s3_image_at"`
								RatingsCount struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr"`
								} `xml:"ratings_count"`
								TextReviewsCount struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr"`
								} `xml:"text_reviews_count"`
								AuthorProgramAt struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr"`
									Nil  string `xml:"nil,attr"`
								} `xml:"author_program_at"`
								BestBookID struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr"`
								} `xml:"best_book_id"`
								SortByName       string `xml:"sort_by_name"`
								ShelfDisplayName string `xml:"shelf_display_name"`
								AuthorURI        string `xml:"author_uri"`
							} `xml:"author"`
						} `xml:"book"`
					} `xml:"user_status"`
				} `xml:"object"`
				Body string `xml:"body"`
			} `xml:"update"`
		} `xml:"updates"`
		UserStatuses struct {
			Text       string `xml:",chardata"`
			UserStatus []struct {
				Text string `xml:",chardata"`
				ID   struct {
					Text string `xml:",chardata"`
					Type string `xml:"type,attr"`
				} `xml:"id"`
				Page struct {
					Text string `xml:",chardata"`
					Type string `xml:"type,attr"`
				} `xml:"page"`
				CommentsCount struct {
					Text string `xml:",chardata"`
					Type string `xml:"type,attr"`
				} `xml:"comments_count"`
				Body      string `xml:"body"`
				CreatedAt struct {
					Text string `xml:",chardata"`
					Type string `xml:"type,attr"`
				} `xml:"created_at"`
				UpdatedAt struct {
					Text string `xml:",chardata"`
					Type string `xml:"type,attr"`
				} `xml:"updated_at"`
				LastCommentAt struct {
					Text string `xml:",chardata"`
					Type string `xml:"type,attr"`
					Nil  string `xml:"nil,attr"`
				} `xml:"last_comment_at"`
				RatingsCount struct {
					Text string `xml:",chardata"`
					Type string `xml:"type,attr"`
				} `xml:"ratings_count"`
				Percent struct {
					Text string `xml:",chardata"`
					Type string `xml:"type,attr"`
				} `xml:"percent"`
				WorkID struct {
					Text string `xml:",chardata"`
					Type string `xml:"type,attr"`
				} `xml:"work_id"`
				NoteURI struct {
					Text string `xml:",chardata"`
					Nil  string `xml:"nil,attr"`
				} `xml:"note_uri"`
				NoteUpdatedAt struct {
					Text string `xml:",chardata"`
					Type string `xml:"type,attr"`
					Nil  string `xml:"nil,attr"`
				} `xml:"note_updated_at"`
				ReviewID struct {
					Text string `xml:",chardata"`
					Type string `xml:"type,attr"`
				} `xml:"review_id"`
				Book struct {
					Text string `xml:",chardata"`
					ID   struct {
						Text string `xml:",chardata"`
						Type string `xml:"type,attr"`
					} `xml:"id"`
					Title    string `xml:"title"`
					NumPages struct {
						Text string `xml:",chardata"`
						Type string `xml:"type,attr"`
					} `xml:"num_pages"`
					ImageURL      string `xml:"image_url"`
					SmallImageURL string `xml:"small_image_url"`
					Authors       struct {
						Text   string `xml:",chardata"`
						Author []struct {
							Text string `xml:",chardata"`
							ID   string `xml:"id"`
							Name string `xml:"name"`
						} `xml:"author"`
					} `xml:"authors"`
				} `xml:"book"`
			} `xml:"user_status"`
		} `xml:"user_statuses"`
	} `xml:"user"`
}

type ReadStatus struct {
	XMLName xml.Name `xml:"GoodreadsResponse"`
	Text    string   `xml:",chardata"`
	Request struct {
		Text           string `xml:",chardata"`
		Authentication string `xml:"authentication"`
		Key            string `xml:"key"`
		Method         string `xml:"method"`
	} `xml:"Request"`
	UserStatus struct {
		Text          string   `xml:",chardata"`
		ID            string   `xml:"id"`
		Header        string   `xml:"header"`
		Body          string   `xml:"body"`
		CreatedAt     string   `xml:"created_at"`
		UpdatedAt     []string `xml:"updated_at"`
		LikesCount    string   `xml:"likes_count"`
		CommentsCount string   `xml:"comments_count"`
		Liked         string   `xml:"liked"`
		Page          string   `xml:"page"`
		Percent       string   `xml:"percent"`
		WorkID        string   `xml:"work_id"`
		User          struct {
			Text          string `xml:",chardata"`
			ID            string `xml:"id"`
			URI           string `xml:"uri"`
			Name          string `xml:"name"`
			DisplayName   string `xml:"display_name"`
			Location      string `xml:"location"`
			Link          string `xml:"link"`
			ImageURL      string `xml:"image_url"`
			SmallImageURL string `xml:"small_image_url"`
			HasImage      string `xml:"has_image"`
		} `xml:"user"`
		Book struct {
			Text               string `xml:",chardata"`
			ID                 string `xml:"id"`
			Title              string `xml:"title"`
			TitleWithoutSeries string `xml:"title_without_series"`
			Link               string `xml:"link"`
			SmallImageURL      string `xml:"small_image_url"`
			ImageURL           string `xml:"image_url"`
			NumPages           string `xml:"num_pages"`
			Work               struct {
				Text string `xml:",chardata"`
				ID   string `xml:"id"`
			} `xml:"work"`
			ISBN             string `xml:"isbn"`
			Isbn13           string `xml:"isbn13"`
			AverageRating    string `xml:"average_rating"`
			RatingsCount     string `xml:"ratings_count"`
			PublicationYear  string `xml:"publication_year"`
			PublicationMonth string `xml:"publication_month"`
			PublicationDay   string `xml:"publication_day"`
			Authors          struct {
				Text   string `xml:",chardata"`
				Author struct {
					Text string `xml:",chardata"`
					ID   string `xml:"id"`
					Name string `xml:"name"`
					Link string `xml:"link"`
				} `xml:"author"`
			} `xml:"authors"`
		} `xml:"book"`
		Comments struct {
			Text  string `xml:",chardata"`
			Start string `xml:"start,attr"`
			End   string `xml:"end,attr"`
			Total string `xml:"total,attr"`
		} `xml:"comments"`
		PreviousUpdates string `xml:"previous_updates"`
		UserStatus      []struct {
			Text          string `xml:",chardata"`
			ID            string `xml:"id"`
			Body          string `xml:"body"`
			CreatedAt     string `xml:"created_at"`
			LikesCount    string `xml:"likes_count"`
			CommentsCount string `xml:"comments_count"`
			Page          string `xml:"page"`
			Percent       string `xml:"percent"`
			UpdatedAt     string `xml:"updated_at"`
			WorkID        string `xml:"work_id"`
		} `xml:"user_status"`
	} `xml:"user_status"`
}
