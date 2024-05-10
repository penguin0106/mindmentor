import * as React from "react";
import "./book.scss"

function BookComponent() {
    return (
        <div className="book-store">
            <div className="book-slide">
                <div className="bookScroll">
                    <div className="book js-flickity" data-flickity-options='{ "wrapAround": true }'>
                        <div className="book-cell">
                            <div className="book-img">
                                <img src="https://images-na.ssl-images-amazon.com/images/I/81WcnNQ-TBL.jpg" alt=""
                                     className="book-photo"/>
                            </div>
                            <div className="book-content">
                                <div className="book-title">BIG MAGIC</div>
                                <div className="book-author">by Elizabeth Gilbert</div>
                                <div className="rate">
                                    <fieldset className="rating">
                                        <input type="checkbox" id="star5" name="rating" value="5"/>
                                        <label className="full" htmlFor="star5"></label>
                                        <input type="checkbox" id="star4" name="rating" value="4"/>
                                        <label className="full" htmlFor="star4"></label>
                                        <input type="checkbox" id="star3" name="rating" value="3"/>
                                        <label className="full" htmlFor="star3"></label>
                                        <input type="checkbox" id="star2" name="rating" value="2"/>
                                        <label className="full" htmlFor="star2"></label>
                                        <input type="checkbox" id="star1" name="rating" value="1"/>
                                        <label className="full" htmlFor="star1"></label>
                                    </fieldset>
                                    <span className="book-voters">1.987 voters</span>
                                </div>
                                <div className="book-sum">Readers of all ages and walks of life have drawn inspiration and
                                    empowerment from Elizabeth Gilbert’s books for years.
                                </div>
                                <div className="book-see">See The Book</div>
                            </div>
                        </div>
                        <div className="book-cell">
                            <div className="book-img">
                                <img src="https://i.pinimg.com/originals/a8/b9/ff/a8b9ff74ed0f3efd97e09a7a0447f892.jpg"
                                     alt="" className="book-photo"/>
                            </div>
                            <div className="book-content">
                                <div className="book-title">Ten Thousand Skies Above You</div>
                                <div className="book-author">by Claudia Gray</div>
                                <div className="rate">
                                    <fieldset className="rating blue">
                                        <input type="checkbox" id="star6" name="rating" value="5"/>
                                        <label className="full1" htmlFor="star6"></label>
                                        <input type="checkbox" id="star7" name="rating" value="4"/>
                                        <label className="full1" htmlFor="star7"></label>
                                        <input type="checkbox" id="star8" name="rating" value="3"/>
                                        <label className="full1" htmlFor="star8"></label>
                                        <input type="checkbox" id="star9" name="rating" value="2"/>
                                        <label className="full1" htmlFor="star9"></label>
                                        <input type="checkbox" id="star10" name="rating" value="1"/>
                                        <label className="full1" htmlFor="star10"></label>
                                    </fieldset>
                                    <span className="book-voters">1.987 voters</span>
                                </div>
                                <div className="book-sum">The hunt for each splinter of Paul's soul sends Marguerite racing
                                    through a war-torn San Francisco.
                                </div>
                                <div className="book-see book-blue">See The Book</div>
                            </div>
                        </div>
                        <div className="book-cell">
                            <div className="book-img">
                                <img src="https://www.abebooks.com/blog/wp-content/uploads/2016/08/10.jpg" alt=""
                                     className="book-photo"/>
                            </div>
                            <div className="book-content">
                                <div className="book-title">A Tale For The Time Being</div>
                                <div className="book-author">by Ruth Ozeki</div>
                                <div className="rate">
                                    <fieldset className="rating purple">
                                        <input type="checkbox" id="star11" name="rating" value="5"/>
                                        <label className="full" htmlFor="star11"></label>
                                        <input type="checkbox" id="star12" name="rating" value="4"/>
                                        <label className="full" htmlFor="star12"></label>
                                        <input type="checkbox" id="star13" name="rating" value="3"/>
                                        <label className="full" htmlFor="star13"></label>
                                        <input type="checkbox" id="star14" name="rating" value="2"/>
                                        <label className="full" htmlFor="star14"></label>
                                        <input type="checkbox" id="star15" name="rating" value="1"/>
                                        <label className="full" htmlFor="star15"></label>
                                    </fieldset>
                                    <span className="book-voters">1.987 voters</span>
                                </div>
                                <div className="book-sum">In Tokyo, sixteen-year-old Nao has decided there’s only one escape
                                    from her aching loneliness and her classmates’ bullying.
                                </div>
                                <div className="book-see book-pink">See The Book</div>
                            </div>
                        </div>
                        <div className="book-cell">
                            <div className="book-img">
                                <img src="https://images-na.ssl-images-amazon.com/images/I/81af+MCATTL.jpg" alt=""
                                     className="book-photo"/>
                            </div>
                            <div className="book-content">
                                <div className="book-title">The Great Gatsby</div>
                                <div className="book-author">by F.Scott Fitzgerald</div>
                                <div className="rate">
                                    <fieldset className="rating yellow">
                                        <input type="checkbox" id="star16" name="rating" value="5"/>
                                        <label className="full" htmlFor="star16"></label>
                                        <input type="checkbox" id="star17" name="rating" value="4"/>
                                        <label className="full" htmlFor="star17"></label>
                                        <input type="checkbox" id="star18" name="rating" value="3"/>
                                        <label className="full" htmlFor="star18"></label>
                                        <input type="checkbox" id="star19" name="rating" value="2"/>
                                        <label className="full" htmlFor="star19"></label>
                                        <input type="checkbox" id="star20" name="rating" value="1"/>
                                        <label className="full" htmlFor="star20"></label>
                                    </fieldset>
                                    <span className="book-voters">1.987 voters</span>
                                </div>
                                <div className="book-sum">The Great Gatsby, F. Scott Fitzgerald’s third book, stands as the
                                    supreme achievement of his career.
                                </div>
                                <div className="book-see book-yellow">See The Book</div>
                            </div>
                        </div>
                        <div className="book-cell">
                            <div className="book-img">
                                <img src="https://images-na.ssl-images-amazon.com/images/I/81UWB7oUZ0L.jpg" alt=""
                                     className="book-photo"/>
                            </div>
                            <div className="book-content">
                                <div className="book-title">After You</div>
                                <div className="book-author">by Jojo Moyes</div>
                                <div className="rate">
                                    <fieldset className="rating dark-purp">
                                        <input type="checkbox" id="star21" name="rating" value="5"/>
                                        <label className="full" htmlFor="star21"></label>
                                        <input type="checkbox" id="star22" name="rating" value="4"/>
                                        <label className="full" htmlFor="star22"></label>
                                        <input type="checkbox" id="star23" name="rating" value="3"/>
                                        <label className="full" htmlFor="star23"></label>
                                        <input type="checkbox" id="star24" name="rating" value="2"/>
                                        <label className="full" htmlFor="star24"></label>
                                        <input type="checkbox" id="star25" name="rating" value="1"/>
                                        <label className="full" htmlFor="star25"></label>
                                    </fieldset>
                                    <span className="book-voters">1.987 voters</span>
                                </div>
                                <div className="book-sum">Louisa Clark is no longer just an ordinary girl living an ordinary
                                    life. After the transformative six months spent.
                                </div>
                                <div className="book-see book-purple">See The Book</div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
            <div className="main-wrapper">
                <div className="books-of">
                    <div className="week">
                        <div className="author-title">Authorof the week</div>
                        {/* More author elements */}
                    </div>
                    <div className="week year">
                        <div className="author-title">Books of the year</div>
                        {/* More year-book elements */}
                    </div>
                    <div className="overlay"></div>
                </div>
                <div className="popular-books">
                    <div className="main-menu">
                        <div className="genre">Popular by Genre</div>
                        <div className="book-types">
                            <a href="#" className="book-type active"> All Genres</a>
                            <a href="#" className="book-type"> Business</a>
                            <a href="#" className="book-type"> Science</a>
                            <a href="#" className="book-type"> Fiction</a>
                            <a href="#" className="book-type"> Philosophy</a>
                            <a href="#" className="book-type"> Biography</a>
                        </div>
                    </div>
                    <div className="book-cards">
                        <div className="book-card">
                            <div className="content-wrapper">
                                <img
                                    src="https://imagesvc.meredithcorp.io/v3/mm/image?url=https%3A%2F%2Fstatic.onecms.io%2Fwp-content%2Fuploads%2Fsites%2F6%2F2019%2F07%2Fchances-are-1-2000.jpg&q=85"
                                    alt="" className="book-card-img"/>
                                <div className="card-content">
                                    <div className="book-name">Changes Are</div>
                                    <div className="book-by">by Richard Russo</div>
                                    <div className="rate">
                                        <fieldset className="rating book-rate">
                                            <input type="checkbox" id="star-c1" name="rating" value="5"/>
                                            <label className="full" htmlFor="star-c1"></label>
                                            <input type="checkbox" id="star-c2" name="rating" value="4"/>
                                            <label className="full" htmlFor="star-c2"></label>
                                            <input type="checkbox" id="star-c3" name="rating" value="3"/>
                                            <label className="full" htmlFor="star-c3"></label>
                                            <input type="checkbox" id="star-c4" name="rating" value="2"/>
                                            <label className="full" htmlFor="star-c4"></label>
                                            <input type="checkbox" id="star-c5" name="rating" value="1"/>
                                            <label className="full" htmlFor="star-c5"></label>
                                        </fieldset>
                                        <span className="book-voters card-vote">1.987 voters</span>
                                    </div>
                                    <div className="book-sum card-sum">
                                        Readers of all ages and walks of life have drawn inspiration and empowerment
                                        from Elizabeth Gilbert’s books for years.
                                    </div>
                                </div>
                            </div>
                            <div className="likes">
                                <div className="like-profile">
                                    <img src="https://randomuser.me/api/portraits/women/63.jpg" alt=""
                                         className="like-img"/>
                                </div>
                                <div className="like-profile">
                                    <img src="https://pbs.twimg.com/profile_images/2452384114/noplz47r59v1uxvyg8ku.png"
                                         alt="" className="like-img"/>
                                </div>
                                <div className="like-profile">
                                    <img
                                        src="https://images.unsplash.com/photo-1535713875002-d1d0cf377fde?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=1400&q=80"
                                        alt="" className="like-img"/>
                                </div>
                                <div className="like-name">
                                    <span>Samantha William</span> and<span> 2 other friends</span> like this
                                </div>
                            </div>
                        </div>
                        <div className="book-card">
                            <div className="content-wrapper">
                                <img
                                    src="https://imagesvc.meredithcorp.io/v3/mm/image?url=https%3A%2F%2Fstatic.onecms.io%2Fwp-content%2Fuploads%2Fsites%2F6%2F2019%2F07%2Fchances-are-1-2000.jpg&q=85"
                                    alt="" className="book-card-img"/>
                                <div className="card-content">
                                    <div className="book-name">Changes Are</div>
                                    <div className="book-by">by Richard Russo</div>
                                    <div className="rate">
                                        <fieldset className="rating book-rate">
                                            <input type="checkbox" id="star-c1" name="rating" value="5"/>
                                            <label className="full" htmlFor="star-c1"></label>
                                            <input type="checkbox" id="star-c2" name="rating" value="4"/>
                                            <label className="full" htmlFor="star-c2"></label>
                                            <input type="checkbox" id="star-c3" name="rating" value="3"/>
                                            <label className="full" htmlFor="star-c3"></label>
                                            <input type="checkbox" id="star-c4" name="rating" value="2"/>
                                            <label className="full" htmlFor="star-c4"></label>
                                            <input type="checkbox" id="star-c5" name="rating" value="1"/>
                                            <label className="full" htmlFor="star-c5"></label>
                                        </fieldset>
                                        <span className="book-voters card-vote">1.987 voters</span>
                                    </div>
                                    <div className="book-sum card-sum">
                                        Readers of all ages and walks of life have drawn inspiration and empowerment
                                        from Elizabeth Gilbert’s books for years.
                                    </div>
                                </div>
                            </div>
                            <div className="likes">
                                <div className="like-profile">
                                    <img src="https://randomuser.me/api/portraits/women/63.jpg" alt=""
                                         className="like-img"/>
                                </div>
                                <div className="like-profile">
                                    <img src="https://pbs.twimg.com/profile_images/2452384114/noplz47r59v1uxvyg8ku.png"
                                         alt="" className="like-img"/>
                                </div>
                                <div className="like-profile">
                                    <img
                                        src="https://images.unsplash.com/photo-1535713875002-d1d0cf377fde?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=1400&q=80"
                                        alt="" className="like-img"/>
                                </div>
                                <div className="like-name">
                                    <span>Samantha William</span> and<span> 2 other friends</span> like this
                                </div>
                            </div>
                        </div>
                        <div className="book-card">
                            <div className="content-wrapper">
                                <img
                                    src="https://imagesvc.meredithcorp.io/v3/mm/image?url=https%3A%2F%2Fstatic.onecms.io%2Fwp-content%2Fuploads%2Fsites%2F6%2F2019%2F07%2Fchances-are-1-2000.jpg&q=85"
                                    alt="" className="book-card-img"/>
                                <div className="card-content">
                                    <div className="book-name">Changes Are</div>
                                    <div className="book-by">by Richard Russo</div>
                                    <div className="rate">
                                        <fieldset className="rating book-rate">
                                            <input type="checkbox" id="star-c1" name="rating" value="5"/>
                                            <label className="full" htmlFor="star-c1"></label>
                                            <input type="checkbox" id="star-c2" name="rating" value="4"/>
                                            <label className="full" htmlFor="star-c2"></label>
                                            <input type="checkbox" id="star-c3" name="rating" value="3"/>
                                            <label className="full" htmlFor="star-c3"></label>
                                            <input type="checkbox" id="star-c4" name="rating" value="2"/>
                                            <label className="full" htmlFor="star-c4"></label>
                                            <input type="checkbox" id="star-c5" name="rating" value="1"/>
                                            <label className="full" htmlFor="star-c5"></label>
                                        </fieldset>
                                        <span className="book-voters card-vote">1.987 voters</span>
                                    </div>
                                    <div className="book-sum card-sum">
                                        Readers of all ages and walks of life have drawn inspiration and empowerment
                                        from Elizabeth Gilbert’s books for years.
                                    </div>
                                </div>
                            </div>
                            <div className="likes">
                                <div className="like-profile">
                                    <img src="https://randomuser.me/api/portraits/women/63.jpg" alt=""
                                         className="like-img"/>
                                </div>
                                <div className="like-profile">
                                    <img src="https://pbs.twimg.com/profile_images/2452384114/noplz47r59v1uxvyg8ku.png"
                                         alt="" className="like-img"/>
                                </div>
                                <div className="like-profile">
                                    <img
                                        src="https://images.unsplash.com/photo-1535713875002-d1d0cf377fde?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=1400&q=80"
                                        alt="" className="like-img"/>
                                </div>
                                <div className="like-name">
                                    <span>Samantha William</span> and<span> 2 other friends</span> like this
                                </div>
                            </div>
                        </div>
                        {/* More book-card elements */}
                    </div>
                </div>
            </div>
        </div>
    );
}

export default BookComponent;


