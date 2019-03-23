
// -*- JavaScript -*-


class MemberItem extends React.Component {
  render() {
    return (
      <tr>
        <td> {this.props.id}    </td>
        <td> {this.props.firstname} </td>
        <td> {this.props.lastname}  </td>
      </tr>
    );
  }
}

class MembersList extends React.Component {
  constructor(props) {
    super(props);
    this.state = { members: [] };
  }

  componentDidMount() {
    this.serverRequest =
      axios
        .get("/members")
        .then((result) => {
            console.log('Received members: ' + result.data);
           this.setState({ members: result.data });
        });
  }

  render() {
    const members = this.state.members.map((member, i) => {
      return (
        <MemberItem key={i} id={member.Id} firstname={member.Firstname} lastname={member.Lastname} />
      );
    });

    return (
      <div>
        <table><tbody>
          <tr><th>Id</th><th>Firstname</th><th>Lastname</th></tr>
          {members}
        </tbody></table>

      </div>
    );
  }
}

ReactDOM.render( <MembersList/>, document.querySelector("#root"));
