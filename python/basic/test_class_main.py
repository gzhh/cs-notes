"""
测试类：
unittest 常用断言方法
1. assertEqual(a, b)
2. assertNotEqual(a, b)
3. assertTrue(x)
4. assertFalse(x)
5. assertIn(item , list)
6. assertNotIn(item , list)

setUp()方法
    通用对象只需创建一次，就能在每个测试方法中使用
"""

import unittest

class AnonymousSurvey:
    """Collect anonymous answers to a survey question."""

    def __init__(self, question):
        """Store a question, and prepare to store responses."""
        self.question = question
        self.responses = []

    def show_question(self):
        """Show the survey question."""
        print(self.question)

    def store_response(self, new_response):
        """Store a single response to the survey."""
        self.responses.append(new_response)

    def show_results(self):
        """Show all the responses that have been given."""
        print("Survey results:")
        for response in self.responses:
            print(f"- {response}")

"""
if __name__ == '__main__':
    # Define a question, and make a survey.
    question = "What language did you first learn to speak?"
    my_survey = AnonymousSurvey(question)

    # Show the question, and store responses to the question.
    my_survey.show_question()
    print("Enter 'q' at any time to quit.\n")
    while True:
        response = input("Language: ")
        if response == 'q':
            break
        my_survey.store_response(response)

    # Show the survey results.
    print("\nThank you to everyone who participated in the survey!")
    my_survey.show_results()
"""


class TestAnonymousSurvey(unittest.TestCase):
    """Tests for the class AnonymousSurvey"""

    def setUp(self):
        """
        Create a survey and a set of responses for use in all test methods.
        """
        question = "What language did you first learn to speak?"
        self.my_survey = AnonymousSurvey(question)
        self.responses = ['English', 'Spanish', 'Mandarin']

    def test_store_single_response(self):
        """Test that a single response is stored properly."""
        self.my_survey.store_response(self.responses[0])
        self.assertIn(self.responses[0], self.my_survey.responses) # responses[0]

    def test_store_three_responses(self):
        """Test that three individual responses are stored properly."""
        for response in self.responses:
            self.my_survey.store_response(response)
        for response in self.responses:
            self.assertIn(response, self.my_survey.responses)

if __name__ == '__main__':
    unittest.main()